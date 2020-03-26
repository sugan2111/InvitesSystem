package services

import "math"

func CalculateDistance(latitudeFrom, longitudeFrom, latitudeTo, longitudeTo float64) float64 {
	long1 := longitudeFrom * (math.Pi / 180)
	long2 := longitudeTo * (math.Pi / 180)
	lat1 := latitudeFrom * (math.Pi / 180)
	lat2 := latitudeTo * (math.Pi / 180)

	//Haversine Formula
	dlong := long2 - long1
	dlati := lat2 - lat1

	val := math.Pow(math.Sin(dlati/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlong/2), 2)

	res := 2 * math.Asin(math.Sqrt(val))

	radius := 6371

	return res * float64(radius)

}

