# pocproduct

A minimum viable product demonstrating Golang, dynamodb, Docker, and various frameworks being used to create a RESTful web service

# Running a local version 

****Note**** 
Does not work with Docker

1. Build the executable on your machine and run it : 
  1. go build (inside the pocproduct directory) 
  2. Create two environmental variables. One for AWS_ACCESS_KEY_ID, and AWS_SECRET_ACCESS_KEY. Link them to an IAM user.
  3. ./pocproduct.exe (inside the pocproduct directory) 

2. Pull down DynamoDB local version using Docker: docker pull dwmkerr/dynamodb

3. Run the local dynamodb version: docker run -p 8000:8000 dwmkerr/dynamodb

4. To create a table called "Products" that the app uses, send a POST statement to localhost:8080/local endpoint

If you use Postman, here is a collections https://www.getpostman.com/collections/39c769ab78f7bec5a9b7

# Connecting to Actual DynamoDb resource

1. Run docker on your machine.

2. Pull down image: docker pull whewitt363/pocproduct in the command line

3. Run image: docker run -e AWS_ACCESS_KEY_ID='actual thing' -e AWS_SECRET_ACCESS_KEY='actual thing' -p 8080:8080 -it --rm --name pocproduct pocproduct

4. Create a "Products" table in your dynamodb application with a Primary key called "Id", that's all.

# Using Endpoint

Look at the Routes.go file if you want an overview of them, otherwise: 

If you use Postman, here is a collections https://www.getpostman.com/collections/86794713779a6956d64e

If not, here are some routes

localhost:8080/local/product/1 POST
  Body: {"price":0.75,"name":"kiwi"}
localhost:8080/local/product/1 GET
localhost:8080/local/product PUT
  Body: {"Id":"1","price":1000,"name":"banana"}
localhost:8080/local/product/1 DELETE

Remote
localhost:8080/aws/product/1 POST
  Body: {"price":0.75,"name":"kiwi"}
localhost:8080/aws/product/1 GET
localhost:8080/aws/product PUT
  Body: {"Id":"1","price":1000,"name":"banana"}
localhost:8080/aws/product/1 DELETE

# Live Version: Coming sooooon!

# Assumptions

1. The ask was for a minimalist project. Use go, dynamo(both local and otherwise) not overkill.

# Future Work

1. Integration Testing, If I had more time, then testing would have been a much larger concern. Integration testing to make sure the dynamodb connection is available, that components in the app are working together, etc. 

2. Unit Testing, If I had more time then unit testing the logic throughout the application would be a must, lots of it would be difficult to update and change later.

3. Refactoring. The handlers are nowhere near pretty. Lots of code reuse, some loops are poorly written, or even missing. 

4. Experiment with DynamoDB libraries. I stuck to the standard AWS libraries for comunicating with Dynamo, but Underarmour's project dynago looks cool, and in general, different libraries for it.

5. A different HTTP library. Mux Gorilla seemed like an easy solution, but it really doesn't do much more than the basic HTTP handling libraries built into go.

6. Expand the querying mechanics. I did not do very well with this part of the assignment, it really should have had more fields and been more expansive. I primarily focused on finishing as opposed to flushing out the capabilities.
