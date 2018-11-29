package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"BikeIndex",
		"GET",
		"/bikes",
		BikeIndex,
	},
	Route{
		"BikeShow",
		"GET",
		"/bikes/{bikeId}",
		BikeShow,
	},
	Route{
		"BikeCreate",
		"POST",
		"/bikes",
		BikeCreate,
	},
}
