package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

func localProductQueryHandler(w http.ResponseWriter, r *http.Request) {

	proj := expression.NamesList(expression.Name("Price"), expression.Name("Id"), expression.Name("Name"))

	//My attempt at using the query map for finding exact values, it is rough
	if len(r.URL.Query()) > 0 {
		filters := make([]expression.ConditionBuilder, len(r.URL.Query()))
		counter := 0
		if len(r.URL.Query()["price"]) > 0 {
			filters[counter] = expression.Name("Price").LessThanEqual(expression.Value(r.URL.Query()["price"]))
			counter++
		}
		if len(r.URL.Query()["name"]) > 0 {
			filters[counter] = expression.Name("Name").Contains(r.URL.Query().Get("name"))
			counter++
		}
		if len(r.URL.Query()["id"]) > 0 {
			filters[counter] = expression.Name("Id").Contains(r.URL.Query().Get("id"))
			counter++
		}

		expr := *new(expression.Expression)
		err := errors.New("")
		switch len(filters) {
		case 1:
			expr, err = expression.NewBuilder().WithFilter(filters[0]).WithProjection(proj).Build()
		case 2:
			expr, err = expression.NewBuilder().WithFilter(filters[0]).WithFilter(filters[1]).WithProjection(proj).Build()
		case 3:
			expr, err = expression.NewBuilder().WithFilter(filters[0]).WithFilter(filters[1]).WithFilter(filters[2]).WithProjection(proj).Build()
		default:
			fmt.Fprintln(w, "invalid parameters! Accepted query parameters are: ('price','name','id') ")
		}

		params := &dynamodb.ScanInput{
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			FilterExpression:          expr.Filter(),
			ProjectionExpression:      expr.Projection(),
			TableName:                 aws.String("Products"),
		}

		// Make the DynamoDB Query API call
		result, err := Localsvc.Scan(params)

		resultProducts := make([]Product, len(result.Items))
		for i, productmap := range result.Items {
			product := Product{}

			err = dynamodbattribute.UnmarshalMap(productmap, &product)

			if err != nil {
				fmt.Println("Error unmarshalling dynamodb response")
				fmt.Println(err.Error())
			}

			resultProducts[i] = product
		}

		productsjson, err := json.Marshal(resultProducts)
		if err != nil {
			fmt.Fprintln(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(productsjson)
		return
	}

	params := &dynamodb.ScanInput{
		TableName: aws.String("Products"),
	}
	result, err := Localsvc.Scan(params)

	resultProducts := make([]Product, len(result.Items))
	for i, productmap := range result.Items {
		product := Product{}

		err = dynamodbattribute.UnmarshalMap(productmap, &product)

		if err != nil {
			fmt.Println("Error unmarshalling dynamodb response")
			fmt.Println(err.Error())
		}

		resultProducts[i] = product
	}

	sort.Slice(resultProducts, func(i, j int) bool {
		return resultProducts[i].Price > resultProducts[j].Price
	})

	productsjson, err := json.Marshal(resultProducts)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(productsjson)

}

func awsProductQueryHandler(w http.ResponseWriter, r *http.Request) {

}
