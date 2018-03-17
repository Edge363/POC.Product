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
		localProductQueryHandler,
	},
	Route{
		"localGet",
		"GET",
		"/local/product/{productid}",
		localProductGetHandler,
	},
	Route{
		"localPut",
		"PUT",
		"/local/product",
		localProductPutHandler,
	},
	Route{
		"localPut",
		"POST",
		"/local/product/{productid}",
		localProductPostHandler,
	},
	Route{
		"Create Dynamo Table",
		"POST",
		"/local",
		createLocalTableDynamo,
	},
	Route{
		"localDelete",
		"DELETE",
		"/local/product/{productid}",
		localProductDeleteHandler,
	},

	//AWS Routes
	Route{
		"awsQuery",
		"GET",
		"/aws/product",
		awsProductQueryHandler,
	},
	Route{
		"awsGet",
		"GET",
		"/aws/product/{productid}",
		awsProductGetHandler,
	},
	Route{
		"awsPut",
		"PUT",
		"/aws/product",
		awsProductPutHandler,
	},
	Route{
		"awsPost",
		"POST",
		"/aws/product/{productid}",
		awsProductPostHandler,
	},
	Route{
		"awsDelete",
		"DELETE",
		"/aws/product/{productid}",
		awsProductDeleteHandler,
	},
}
