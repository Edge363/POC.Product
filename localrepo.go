package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var localsvc *dynamodb.DynamoDB

func init() {

	creds := credentials.NewStaticCredentials("1", "1", "")

	sess, err := session.NewSession(
		&aws.Config{
			Region:      aws.String("us-east-1"),
			Credentials: creds,
			Endpoint:    aws.String("localhost:8000"),
		},
	)
	if err != nil {
		fmt.Println("Error Creating connection with local dynamo")
		os.Exit(1)
	}

	// Create DynamoDB client
	localsvc = dynamodb.New(sess)

}

func getProduct(id *string) {
	result, err := localsvc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Products"),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: id,
			},
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	product := Product{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &product)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

}

func deleteProduct(id *string) {
	_, err := localsvc.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String("Products"),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: id,
			},
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

}

func putProduct(id *string, product *Product) {
	_, errput := localsvc.UpdateItem(&dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: id,
			},
			"Price": {
				N: aws.String(strconv.FormatFloat(product.Price, 'f', -1, 64)),
			},
			"Name": {
				S: aws.String(product.Name),
			},
		},
		TableName: aws.String("Products"),
	})
	if errput != nil {
		fmt.Println(errput.Error())
		return
	}

}

func postProduct(id *string, product *Product) {
	av, errmarshal := dynamodbattribute.MarshalMap(product)
	if errmarshal != nil {
		fmt.Println(errmarshal.Error())
		return
	}
	_, errput := localsvc.PutItem(&dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Products"),
	})
	if errput != nil {
		fmt.Println(errput.Error())
		return
	}

}
