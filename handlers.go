package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func BikeIndex(w http.ResponseWriter, r *http.Request) {

	//bikes := Bikes{
	//	Bike{Desc: "Haro Midway", Frame: "3 tubes Crmo, Internal HT & mid BB - 20.5 & 21 TT", Gearing: "25/9", CustomerPrice: money.Money{CurrencyCode:"CHF", Units:459, Nanos:0}},
	//	Bike{Desc: "Haro CK AM 2019", Frame: "3 tubes Crmo, Internal HT & mid BB - 20.5 & 21 TT", Gearing: "25/9", CustomerPrice: money.Money{CurrencyCode:"CHF", Units:759, Nanos:0}},
	//}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(bikes); err != nil {
		panic(err)
	}

}

func BikeShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bikeId := vars["bikeId"]
	i, _ := strconv.Atoi(bikeId)
	bike := RepoFindBike(i)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(bike); err != nil {
		panic(err)
	}
}

func BikeCreate(w http.ResponseWriter, r *http.Request) {
	var bike Bike
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &bike); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateBike(bike)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}

}
