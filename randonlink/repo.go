package main

 import (
 	"errors"
 	"fmt"
 	"os"
 	"strconv"

 	"github.com/aws/aws-sdk-go/aws/credentials"
 	"github.com/aws/aws-sdk-go/aws"
 	"github.com/aws/aws-sdk-go/aws/session"
 	"github.com/aws/aws-sdk-go/service/dynamodb"
 	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
 )

 var Awssvc *dynamodb.DynamoDB

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

 	 Awssvc = dynamodb.New(sess)

 }

 func getProduct(id *string, svc *dynamodb.DynamoDB) (Product, error) {
 	result, err := svc.GetItem(&dynamodb.GetItemInput{
 		TableName: aws.String("Products"),
 		Key: map[string]*dynamodb.AttributeValue{
 			"Id": {
 				S: id,
 			},
 		},
 	})

 	if err != nil {
 		fmt.Println(err.Error())
 		return *new(Product), err
 	}

 	product := Product{}

 	err = dynamodbattribute.UnmarshalMap(result.Item, &product)

 	if err != nil {
 		fmt.Println(err.Error())
 		return *new(Product), err
 	}
 	return product, nil
 }

 func deleteProduct(id *string, svc *dynamodb.DynamoDB) error {
 	_, err := svc.DeleteItem(&dynamodb.DeleteItemInput{
 		TableName: aws.String("Products"),
 		Key: map[string]*dynamodb.AttributeValue{
 			"Id": {
 				S: id,
 			},
 		},
 	})

 	if err != nil {
 		fmt.Println(err.Error())
 		return err
 	}

 	return nil
 }

 func putProduct(product *Product, svc *dynamodb.DynamoDB) error {
 	_, errput := svc.PutItem(&dynamodb.PutItemInput{
 		Item: map[string]*dynamodb.AttributeValue{
 			"Id": {
 				S: &product.Id,
 			},
 			"Name": {
 				S: aws.String(product.Name),
 			},
 			"Price": {
 				N: aws.String(strconv.FormatFloat(product.Price, 'f', -1, 64)),
 			},
 		},
 		TableName: aws.String("Products"),
 	})
 	if errput != nil {
 		fmt.Println(errput.Error())
 		return errput
 	}

 	return nil
 }

 func postProduct(id *string, product *Product, svc *dynamodb.DynamoDB) error {
 	productCheck, err := getProduct(id, svc)
 	if err != nil {
 		return err
 	}
 	if productCheck.Id != "" {
 		return errors.New("resource already exists, use put to replace existing resources")
 	}
 	product.Id = *id
 	putProduct(product, svc)
 	return nil
 }
