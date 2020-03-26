package models

import (
	"context"
	"encoding/json"
	"fmt"
	helper "github.com/sugan2111/InvitesSystem/helpers"
	"github.com/sugan2111/InvitesSystem/store"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStore struct {
	*mongo.Collection
}

func NewClient(uri string) MongoStore {
	return MongoStore{helper.ConnectDB()}

}

func (d MongoStore) Insert(customer store.Customer, w http.ResponseWriter) error {
	collection := helper.ConnectDB()

	result, err := collection.InsertOne(context.TODO(), customer)

	if err != nil {
		return fmt.Errorf("unable to insert item:%v", err)
	}
	json.NewEncoder(w).Encode(result)
	return nil
}
