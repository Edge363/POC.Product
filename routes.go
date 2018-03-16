package main

import (
	"net/http"
)

//Route is a struct for making the routes more easily viewed
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes is a supporter type for making the neato default routes object
type Routes []Route

var routes = Routes{
	//Local Routes
	Route{
		"Index",
		"GET",
		"/local/product",
		localProductQueryHandler,
	},
	Route{
		"Index",
		"GET",
		"/local/product/{productid}",
		localProductGetHandler,
	},
	Route{
		"TodoIndex",
		"PUT",
		"/local/product/{productid}",
		localProductPutHandler,
	},
	Route{
		"TodoShow",
		"POST",
		"local/product/{productid}",
		localProductPostHandler,
	},
	Route{
		"TodoShow",
		"DELETE",
		"local/product/{productid}",
		localProductDeleteHandler,
	},

	//AWS Routes
	Route{
		"Index",
		"GET",
		"/aws/product",
		awsProductQueryHandler,
	},
	Route{
		"Index",
		"GET",
		"/aws/product/{productid}",
		awsProductGetHandler,
	},
	Route{
		"TodoIndex",
		"PUT",
		"/aws/product/{productid}",
		awsProductPutHandler,
	},
	Route{
		"TodoShow",
		"POST",
		"aws/product/{productid}",
		awsProductPostHandler,
	},
	Route{
		"TodoShow",
		"DELETE",
		"aws/product/{productid}",
		awsProductDeleteHandler,
	},
}
