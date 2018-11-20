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
		"helloWorld",
		"GET",
		"/hello",
		helloHandler,
	},
	// Route{
	// 	"localQuery",
	// 	"GET",
	// 	"/local/product",
	// 	productQueryHandler,
	// },
	// Route{
	// 	"localGet",
	// 	"GET",
	// 	"/local/product/{productid}",
	// 	getHandler,
	// },
	// Route{
	// 	"localPut",
	// 	"PUT",
	// 	"/local/product",
	// 	putHandler,
	// },
	// Route{
	// 	"localPut",
	// 	"POST",
	// 	"/local/product/{productid}",
	// 	postHandler,
	// },
	// Route{
	// 	"Create Dynamo Table",
	// 	"POST",
	// 	"/local",
	// 	postHandler,
	// },
	// Route{
	// 	"localDelete",
	// 	"DELETE",
	// 	"/local/product/{productid}",
	// 	deleteHandler,
	// },
}
