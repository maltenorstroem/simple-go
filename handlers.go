package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// IANA konforme Links /
// List of official link rels:
// http://www.iana.org/assignments/link-relations/link-relations.xhtml
type Link struct {
	Rel string `json:"rel,omitempty"`
	// HTTP Verb
	Method string `json:"method,omitempty"`
	// Absolute URI
	Href string `json:"href,omitempty"`
	// Der mime type des Links.
	Type string `json:"type,omitempty"`
}

type EntityStruct struct {
	Data  *Bike   `json:"data,omitempty"`
	Links []*Link `json:"links,omitempty"`
}

type CollectionStruct struct {
	Data  []*EntityStruct `json:"data,omitempty"`
	Links []*Link         `json:"links,omitempty"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func BikeIndex(w http.ResponseWriter, r *http.Request) {

	var data []*EntityStruct

	for i := 0; i < len(bikes); i++ {
		data = append(data, &EntityStruct{Data: &Bike{Id: bikes[i].Id, Desc: bikes[i].Desc, Frame: bikes[i].Frame, Gearing: bikes[i].Gearing, CustomerPrice: bikes[i].CustomerPrice, SoldOut: bikes[i].SoldOut}})
	}

	collection := CollectionStruct{Data: data}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(collection); err != nil {
		log.Fatal(err)
	}

}

func BikeShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	params := strings.Split(vars["bikeId"], ":")
	bikeId := params[0]
	i, _ := strconv.Atoi(bikeId)
	bike, err := RepoFindBike(i)

	if err == nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		//TODO Mapping funcs definieren: Data: MapTaskToProtoTask(&item), Links: GenerateEntityHateoas(item.Id.String()).Links}
		entity := EntityStruct{Data: &Bike{Id: bike.Id, Desc: bike.Desc, Frame: bike.Frame, Gearing: bike.Gearing, CustomerPrice: bike.CustomerPrice, SoldOut: bike.SoldOut}}
		if err := json.NewEncoder(w).Encode(entity); err != nil {
			log.Fatal(err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		log.Println(err, bikeId)
	}

}

func BikeDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bikeId := vars["bikeId"]
	i, _ := strconv.Atoi(bikeId)
	err := RepoDestroyBike(i)

	if err == nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusAccepted)

	} else {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err)
	}
}

func BikeSetSoldOut(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bikeId := vars["bikeId"]
	i, _ := strconv.Atoi(bikeId)
	bike, _ := RepoFindBike(i)

	bike.SoldOut = true

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusAccepted)

	//TODO Mapping funcs definieren: Data: MapTaskToProtoTask(&item), Links: GenerateEntityHateoas(item.Id.String()).Links}
	entity := EntityStruct{Data: &Bike{Id: bike.Id, Desc: bike.Desc, Frame: bike.Frame, Gearing: bike.Gearing, CustomerPrice: bike.CustomerPrice, SoldOut: bike.SoldOut}}

	if err := json.NewEncoder(w).Encode(entity); err != nil {
		log.Fatal(err)
	}

}

func BikeCreate(w http.ResponseWriter, r *http.Request) {
	var bike Bike
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatal(err)
	}
	if err := r.Body.Close(); err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(body, &bike); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatal(err)
		}
	}

	t := RepoCreateBike(bike)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		log.Fatal(err)
	}

}
