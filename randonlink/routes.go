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
		"helloWorld",
		"GET",
		"/",
		hello,
	},
	Route{
		"productQuery",
		"GET",
		"/product",
		productQuery,
	},
	Route{
		"productGet",
		"GET",
		"/product/{productid}",
		productGet,
	},
	Route{
		"productPut",
		"PUT",
		"/product",
		productPut,
	},
	Route{
		"productPut",
		"POST",
		"/product/{productid}",
		productPost,
	},
	Route{
		"Create Dynamo Table",
		"POST",
		"/productDynamo",
		tablePost,
	},
	Route{
		"productDelete",
		"DELETE",
		"/product/{productid}",
		productDelete,
	},
}
