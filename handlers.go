package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gorilla/mux"

	"encoding/json"
)

func localProductGetHandler(w http.ResponseWriter, r *http.Request) {
	productid := mux.Vars(r)["productid"]
	product, err := getProduct(&productid, Localsvc)
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

func awsProductGetHandler(w http.ResponseWriter, r *http.Request) {
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

func localProductPutHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	product := new(Product)
	err := decoder.Decode(&product)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	err = putProduct(product, Localsvc)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func awsProductPutHandler(w http.ResponseWriter, r *http.Request) {
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

func localProductPostHandler(w http.ResponseWriter, r *http.Request) {
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

func awsProductPostHandler(w http.ResponseWriter, r *http.Request) {
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

func localProductDeleteHandler(w http.ResponseWriter, r *http.Request) {
	productid := mux.Vars(r)["productid"]
	err := deleteProduct(&productid, Localsvc)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func awsProductDeleteHandler(w http.ResponseWriter, r *http.Request) {
	productid := mux.Vars(r)["productid"]
	err := deleteProduct(&productid, Awssvc)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func createLocalTableDynamo(w http.ResponseWriter, r *http.Request) {
	input := &dynamodb.CreateTableInput{
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Id"),
				KeyType:       aws.String("HASH"),
			},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Id"),
				AttributeType: aws.String("S"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String("Products"),
	}

	_, err := Localsvc.CreateTable(input)

	if err != nil {
		fmt.Println("Got error calling CreateTable for local Dynamo")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Created the table Products successfully")
}
