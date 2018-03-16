package main

import (
	"net/http"
)

func localProductGetHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(todos)
}

func awsProductGetHandler(w http.ResponseWriter, r *http.Request) {

}

func localProductQueryHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(todos)
}

func awsProductQueryHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(todos)
}

func localProductPutHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(todos)
}

func awsProductPutHandler(w http.ResponseWriter, r *http.Request) {

}

func localProductPostHandler(w http.ResponseWriter, r *http.Request) {

}

func awsProductPostHandler(w http.ResponseWriter, r *http.Request) {

}

func localProductDeleteHandler(w http.ResponseWriter, r *http.Request) {

}

func awsProductDeleteHandler(w http.ResponseWriter, r *http.Request) {

}
