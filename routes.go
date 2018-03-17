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

//AppRoutes is thhe container for route definitions used by the application
var routes = Routes{
	//Local Routes
	Route{
		"localQuery",
		"GET",
		"/local/product",
		LocalProductQueryHandler,
	},
	Route{
		"localGet",
		"GET",
		"/local/product/{productid}",
		LocalProductGetHandler,
	},
	Route{
		"localPut",
		"PUT",
		"/local/product",
		LocalProductPutHandler,
	},
	Route{
		"localPut",
		"POST",
		"/local/product/{productid}",
		LocalProductPostHandler,
	},
	Route{
		"localDelete",
		"DELETE",
		"/local/product/{productid}",
		LocalProductDeleteHandler,
	},

	//AWS Routes
	Route{
		"awsQuery",
		"GET",
		"/aws/product",
		AwsProductQueryHandler,
	},
	Route{
		"awsGet",
		"GET",
		"/aws/product/{productid}",
		AwsProductGetHandler,
	},
	Route{
		"awsPut",
		"PUT",
		"/aws/product",
		AwsProductPutHandler,
	},
	Route{
		"awsPost",
		"POST",
		"/aws/product/{productid}",
		AwsProductPostHandler,
	},
	Route{
		"awsDelete",
		"DELETE",
		"/aws/product/{productid}",
		AwsProductDeleteHandler,
	},
}
