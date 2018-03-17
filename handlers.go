package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"encoding/json"
)

func LocalProductGetHandler(w http.ResponseWriter, r *http.Request) {

	productid := mux.Vars(r)["productid"]
	product, err := getProduct(&productid, Localsvc)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	productjson, err := json.Marshal(product)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(productjson)
}

func AwsProductGetHandler(w http.ResponseWriter, r *http.Request) {
	productid := mux.Vars(r)["productid"]
	product, err := getProduct(&productid, Awssvc)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	if product.Id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	productjson, err := json.Marshal(product)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(productjson)
}

func LocalProductQueryHandler(w http.ResponseWriter, r *http.Request) {

}

func AwsProductQueryHandler(w http.ResponseWriter, r *http.Request) {

}

func LocalProductPutHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	product := new(Product)
	err := decoder.Decode(&product)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	err = putProduct(product, Awssvc)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func AwsProductPutHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	product := new(Product)
	err := decoder.Decode(&product)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	err = putProduct(product, Awssvc)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func LocalProductPostHandler(w http.ResponseWriter, r *http.Request) {
	productid := mux.Vars(r)["productid"]
	decoder := json.NewDecoder(r.Body)
	product := new(Product)
	err := decoder.Decode(&product)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	err = postProduct(&productid, product, Localsvc)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func AwsProductPostHandler(w http.ResponseWriter, r *http.Request) {
	productid := mux.Vars(r)["productid"]
	decoder := json.NewDecoder(r.Body)
	product := new(Product)
	err := decoder.Decode(&product)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	err = postProduct(&productid, product, Awssvc)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func LocalProductDeleteHandler(w http.ResponseWriter, r *http.Request) {
	productid := mux.Vars(r)["productid"]
	product, err := getProduct(&productid, Localsvc)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	productjson, err := json.Marshal(product)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	w.Write(productjson)
}

func AwsProductDeleteHandler(w http.ResponseWriter, r *http.Request) {
	productid := mux.Vars(r)["productid"]
	err := deleteProduct(&productid, Awssvc)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
