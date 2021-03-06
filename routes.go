package truckly

import "net/http"

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
		"/api/trucks/",
		Index,
	},
	Route{
		"NewTruck",
		"POST",
		"/api/trucks/",
		NewTruck,
	},
	Route{
		"ImportTrucks",
		"POST",
		"/api/import/",
		ImportTrucks,
	},
}
