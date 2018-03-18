# pocproduct

A minimum viable product demonstrating Golang, dynamodb, Docker, and various frameworks being used to create a RESTful web service

# Running a local version

1. Build the executable on your machine and run it: 
  1. go build (inside the pocproduct directory) 
  2. ./pocproduct.exe (inside the pocproduct directory) 

2. Pull down DynamoDB local version using Docker: docker pull dwmkerr/dynamodb

3. Run the local dynamodb version: docker run -p 8000:8000 dwmkerr/dynamodb

4. To create a table called "Products" that the app uses, send a POST statement to localhost:8080/local endpoint

If you use Postman, here is a collections https://www.getpostman.com/collections/39c769ab78f7bec5a9b7

# Connecting to Actual DynamoDb resource

1. Run docker on your machine.

2. Pull down image: docker pull whewitt363/pocproduct in the command line

3. Run image: docker run -e AWS_ACCESS_KEY_ID='actual thing' -e AWS_SECRET_ACCESS_KEY='actual thing' -p 8080:8080 -it --rm --name pocproduct pocproduct

# Using Endpoint

Look at the Routes.go file if you want an overview of them, otherwise: 

If you use Postman, here is a collections https://www.getpostman.com/collections/86794713779a6956d64e

If not, here are some routes

# Local
localhost:8080/local/product/1 POST
  Body: {"price":0.75,"name":"kiwi"}
localhost:8080/local/product/1 GET
localhost:8080/local/product PUT
  Body: {"Id":"1","price":1000,"name":"banana"}
localhost:8080/local/product/1 DELETE

# Remote
localhost:8080/aws/product/1 POST
  Body: {"price":0.75,"name":"kiwi"}
localhost:8080/aws/product/1 GET
localhost:8080/aws/product PUT
  Body: {"Id":"1","price":1000,"name":"banana"}
localhost:8080/aws/product/1 DELETE

# Live Version: Coming sooooon!
