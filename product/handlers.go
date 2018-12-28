package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"encoding/json"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func productGet(w http.ResponseWriter, r *http.Request) {
	productid := mux.Vars(r)["productid"]
	svc := Awssvc

	product, err := getProduct(&productid, svc)
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

func productQuery(w http.ResponseWriter, r *http.Request) {
	productid := mux.Vars(r)["productid"]
	svc := Awssvc

	product, err := getProduct(&productid, svc)
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

func productPut(w http.ResponseWriter, r *http.Request) {
	svc := Awssvc

	decoder := json.NewDecoder(r.Body)
	product := new(Product)
	err := decoder.Decode(&product)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	err = putProduct(product, svc)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func productPost(w http.ResponseWriter, r *http.Request) {
	productid := mux.Vars(r)["productid"]
	svc := Awssvc
	decoder := json.NewDecoder(r.Body)
	product := new(Product)
	err := decoder.Decode(&product)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	err = postProduct(&productid, product, svc)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func productDelete(w http.ResponseWriter, r *http.Request) {
	productid := mux.Vars(r)["productid"]
	svc := Awssvc
	err := deleteProduct(&productid, svc)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
