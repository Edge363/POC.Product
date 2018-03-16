package main

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var awssvc *dynamodb.DynamoDB

func init() {

	// creds := credentials.NewEnvCredentials()

	// sess, err := session.NewSession(&aws.Config{
	// 	Region:      aws.String("us-east-1"),
	// 	Credentials: creds,
	// })

	// if err != nil {
	// 	fmt.Println("Error Creating Session with AWS")
	// 	os.Exit(1)
	// }

	// awssvc = dynamodb.New(sess)

}
