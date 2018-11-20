package main

// import (
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"net/http"
// 	"sort"
// 	"strings"

// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/service/dynamodb"
// 	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
// 	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
// )

// func productQueryHandler(w http.ResponseWriter, r *http.Request) {
// 	svc := Localsvc
// 	if strings.Contains(r.URL.String(), "aws") {
// 		svc = Awssvc
// 	}
// 	if len(r.URL.Query()) > 0 {
// 		queryDynamo(w, r, svc)
// 	}
// 	defaultResponse(w, r, svc)
// }

// func hasParam(r *http.Request, s string) bool {
// 	return len(r.URL.Query()["id"]) > 0
// }

// func addFilters(filters []expression.ConditionBuilder) (expression.Expression, error) {
// 	proj := expression.NamesList(expression.Name("Price"), expression.Name("Id"), expression.Name("Name"))
// 	switch len(filters) {
// 	case 1:
// 		return expression.NewBuilder().WithFilter(filters[0]).WithProjection(proj).Build()
// 	case 2:
// 		return expression.NewBuilder().WithFilter(filters[0]).WithFilter(filters[1]).WithProjection(proj).Build()
// 	default:
// 		return expression.NewBuilder().WithFilter(filters[0]).WithFilter(filters[1]).WithFilter(filters[2]).WithProjection(proj).Build()
// 	}
// }

// func unmarshalDynamoResponse(result *dynamodb.ScanOutput) ([]Product, error) {
// 	resultProducts := make([]Product, len(result.Items))
// 	for i, productmap := range result.Items {
// 		product := Product{}

// 		err := dynamodbattribute.UnmarshalMap(productmap, &product)

// 		if err != nil {
// 			return nil, errors.New("Error unmarshaling dynamo response")
// 		}

// 		resultProducts[i] = product
// 	}

// 	return resultProducts, nil

// }

// func defaultResponse(w http.ResponseWriter, r *http.Request, svc *dynamodb.DynamoDB) {
// 	params := &dynamodb.ScanInput{
// 		TableName: aws.String("Products"),
// 	}
// 	result, err := svc.Scan(params)
// 	if err != nil {
// 		fmt.Fprintln(w, err.Error())
// 		return
// 	}
// 	resultProducts, err := unmarshalDynamoResponse(result)
// 	if err != nil {
// 		fmt.Fprintln(w, err.Error())
// 		return
// 	}
// 	sort.Slice(resultProducts, func(i, j int) bool {
// 		return resultProducts[i].Price > resultProducts[j].Price
// 	})

// 	productsjson, err := json.Marshal(resultProducts)
// 	if err != nil {
// 		fmt.Fprintln(w, err.Error())
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(productsjson)
// }

// func queryDynamo(w http.ResponseWriter, r *http.Request, svc *dynamodb.DynamoDB) {
// 	filters := make([]expression.ConditionBuilder, len(r.URL.Query()))
// 	counter := 0

// 	if hasParam(r, "price") {
// 		filters[counter] = expression.Name("Price").LessThanEqual(expression.Value(r.URL.Query()["price"]))
// 		counter++
// 	}
// 	if hasParam(r, "name") {
// 		filters[counter] = expression.Name("Name").Contains(r.URL.Query().Get("name"))
// 		counter++
// 	}
// 	if hasParam(r, "id") {
// 		filters[counter] = expression.Name("Id").Contains(r.URL.Query().Get("id"))
// 		counter++
// 	}

// 	if counter < len(r.URL.Query()) {
// 		fmt.Fprintln(w, "invalid parameters! Accepted query parameters are: ('price','name','id') ")
// 		return
// 	}
// 	expr, err := addFilters(filters)

// 	params := &dynamodb.ScanInput{
// 		ExpressionAttributeNames:  expr.Names(),
// 		ExpressionAttributeValues: expr.Values(),
// 		FilterExpression:          expr.Filter(),
// 		ProjectionExpression:      expr.Projection(),
// 		TableName:                 aws.String("Products"),
// 	}

// 	// Make the DynamoDB Query API call
// 	result, err := svc.Scan(params)

// 	products, err := unmarshalDynamoResponse(result)
// 	if err != nil {
// 		fmt.Fprintln(w, err.Error())
// 		return
// 	}
// 	productsjson, err := json.Marshal(products)
// 	if err != nil {
// 		fmt.Fprintln(w, err.Error())
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(productsjson)
// 	return
// }
