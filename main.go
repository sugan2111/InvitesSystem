package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/sugan2111/Intercom/services"
	helper "github.com/sugan2111/InvitesSystem/helpers"
	"github.com/sugan2111/InvitesSystem/models"
	"github.com/sugan2111/InvitesSystem/store"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	// better to read this from environment variables
	dbURI = "mongodb://localhost:27017"
)

type Router struct {
	*mux.Router
	store CustomerStore
}

type myData struct {
	Url string `json:"url"`
}

type CustomerStore interface {
	Insert(customer store.Customer, w http.ResponseWriter) error
}

func NewRouter(r *mux.Router, store CustomerStore) Router {
	return Router{Router: r, store: store}
}

func UrlToLines(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return LinesFromReader(resp.Body)
}

func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func (h *Router) storeCustData(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var data myData
	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&data)

	custData, err := UrlToLines(data.Url)
	if err != nil {
		fmt.Println("error:", err)
	}

	h.insertCustData(custData, w)

}

func (h *Router) insertCustData(custData []string, response http.ResponseWriter) {
	var customer store.Customer

	for _, eachline := range custData {
		err := json.Unmarshal([]byte(eachline), &customer)
		if err != nil {
			fmt.Println("error:", err)
		}

		err = h.store.Insert(customer, response)

		if err != nil {
			helper.GetError(err, response)
			return
		}

		custLatitude, err := strconv.ParseFloat(customer.Latitude, 64)
		if err != nil {
			fmt.Println("error:", err)
		}

		custLongitude, err := strconv.ParseFloat(customer.Longitude, 64)
		if err != nil {
			fmt.Println("error:", err)
		}

		distance := services.CalculateDistance(53.339428, -6.257664, custLatitude, custLongitude)
		getInvitees(distance, customer)
	}
}

func getInvitees(distance float64, customer store.Customer) {
	if distance <= 100.0 {
		m := make(map[int]string)

		m[customer.UserId] = customer.Name // Add a new key-value pair

		// To store the keys in slice in sorted order
		var keys []int
		for k := range m {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		// To perform the opertion you want
		for _, k := range keys {
			fmt.Println(k, m[k])
		}
	}
}


func main() {
	fmt.Println("Building the invite management system...")
	store := models.NewClient(dbURI)
	r := NewRouter(mux.NewRouter(), store)
	r.HandleFunc("/customers", r.storeCustData).Methods("POST")
	log.Fatal(http.ListenAndServe(":8001", r))
}
