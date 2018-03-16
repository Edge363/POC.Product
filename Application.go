package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gorilla/mux"
)

var svc *dynamodb.DynamoDB

func init() {

	creds := credentials.NewEnvCredentials()

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: creds,
	})

	if err != nil {
		fmt.Println("Error Creating Session with AWS")
		os.Exit(1)
	}

	svc = dynamodb.New(sess)

}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/local/product", localProductGetHandler).Methods("GET").Queries("filter", "{filter}")
	router.HandleFunc("/local/product", localProductGetHandler).Methods("GET")

	router.HandleFunc("/local/product/{productid}", localProductGetHandler).Methods("GET")
	router.HandleFunc("/aws/product/{productid}", awsProductGetHandler).Methods("GET")

	router.HandleFunc("/local/product/{productid}", localProductPutHandler).Methods("PUT")
	router.HandleFunc("/aws/product/{productid}", awsProductPutHandler).Methods("PUT")

	router.HandleFunc("/local/product/{productid}", localProductPostHandler).Methods("POST")
	router.HandleFunc("/aws/product/{productid}", awsProductPostHandler).Methods("POST")

	router.HandleFunc("/local/product/{productid}", localProductDeleteHandler).Methods("DELETE")
	router.HandleFunc("/aws/product/{productid}", awsProductDeleteHandler).Methods("DELETE")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}

func localProductGetHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "We got it boss")
}

func awsProductGetHandler(w http.ResponseWriter, r *http.Request) {

}

func localProductPutHandler(w http.ResponseWriter, r *http.Request) {

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
