package main

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// 	"strconv"

// 	"github.com/aws/aws-sdk-go/aws/credentials"
// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/dynamodb"
// 	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
// )

// //Localsvc is the connection to a locally running Dynamodb server on port 8000.
// var Localsvc *dynamodb.DynamoDB

// //Awssvc is the connection to your Dynamodb instance running on AWS. This is easily broken if you aren't setting it up correctly.
// var Awssvc *dynamodb.DynamoDB

// func init() {

// 	//create connection to local dynamodb

// 	creds := credentials.NewStaticCredentials("1", "1", "")
// 	sess, err := session.NewSession(
// 		&aws.Config{
// 			Region:   aws.String("us-east-1"),
// 			Endpoint: aws.String("http://localhost:8000"),
// 			Credentials: creds,
// 		},
// 	)
// 	if err != nil {
// 		fmt.Println("Error Creating connection with local dynamo")
// 		os.Exit(1)
// 	}

// 	Localsvc = dynamodb.New(sess)

// 	result, err := Localsvc.ListTables(&dynamodb.ListTablesInput{})

// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	createTable("Products", result)

// 	// Create connection to remote dynamodb
// 	//EXAMPLE
// 	// creds = credentials.NewEnvCredentials()

// 	// sess, err = session.NewSession(&aws.Config{
// 	// 	Region:      aws.String("us-east-1"),
// 	// 	Credentials: creds,
// 	// })

// 	// if err != nil {
// 	// 	fmt.Println("Error Creating Session with AWS")
// 	// 	os.Exit(1)
// 	// }

// 	// Awssvc = dynamodb.New(sess)

// }

// func getProduct(id *string, svc *dynamodb.DynamoDB) (Product, error) {
// 	result, err := svc.GetItem(&dynamodb.GetItemInput{
// 		TableName: aws.String("Products"),
// 		Key: map[string]*dynamodb.AttributeValue{
// 			"Id": {
// 				S: id,
// 			},
// 		},
// 	})

// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return *new(Product), err
// 	}

// 	product := Product{}

// 	err = dynamodbattribute.UnmarshalMap(result.Item, &product)

// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return *new(Product), err
// 	}
// 	return product, nil
// }

// func deleteProduct(id *string, svc *dynamodb.DynamoDB) error {
// 	_, err := svc.DeleteItem(&dynamodb.DeleteItemInput{
// 		TableName: aws.String("Products"),
// 		Key: map[string]*dynamodb.AttributeValue{
// 			"Id": {
// 				S: id,
// 			},
// 		},
// 	})

// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return err
// 	}

// 	return nil
// }

// func putProduct(product *Product, svc *dynamodb.DynamoDB) error {
// 	_, errput := svc.PutItem(&dynamodb.PutItemInput{
// 		Item: map[string]*dynamodb.AttributeValue{
// 			"Id": {
// 				S: &product.Id,
// 			},
// 			"Name": {
// 				S: aws.String(product.Name),
// 			},
// 			"Price": {
// 				N: aws.String(strconv.FormatFloat(product.Price, 'f', -1, 64)),
// 			},
// 		},
// 		TableName: aws.String("Products"),
// 	})
// 	if errput != nil {
// 		fmt.Println(errput.Error())
// 		return errput
// 	}

// 	return nil
// }

// func postProduct(id *string, product *Product, svc *dynamodb.DynamoDB) error {
// 	//check to see if product exists
// 	productCheck, err := getProduct(id, svc)
// 	if err != nil {
// 		return err
// 	}
// 	if productCheck.Id != "" {
// 		return errors.New("resource already exists, use put to replace existing resources")
// 	}
// 	product.Id = *id
// 	putProduct(product, svc)
// 	return nil
// }

// func hasTable(s []*string, e string) bool {
// 	for _, a := range s {
// 		if *a == e {
// 			return true
// 		}
// 	}
// 	return false
// }

// func createTable(name string, result *dynamodb.ListTablesOutput) {
// 	if !hasTable(result.TableNames, name) {
// 		input := &dynamodb.CreateTableInput{
// 			AttributeDefinitions: []*dynamodb.AttributeDefinition{
// 				{
// 					AttributeName: aws.String("Id"),
// 					AttributeType: aws.String("S"),
// 				},
// 			},
// 			KeySchema: []*dynamodb.KeySchemaElement{
// 				{
// 					AttributeName: aws.String("Id"),
// 					KeyType:       aws.String("HASH"),
// 				},
// 			},
// 			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
// 				ReadCapacityUnits:  aws.Int64(10),
// 				WriteCapacityUnits: aws.Int64(10),
// 			},
// 			TableName: aws.String(name),
// 		}
// 		_, err := Localsvc.CreateTable(input)

// 		if err != nil {
// 			fmt.Println("Got error calling CreateTable:")
// 			fmt.Println(err.Error())
// 			os.Exit(1)
// 		}

// 		fmt.Println("Created the table Products")
// 	}
// }
