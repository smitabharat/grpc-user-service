# gRPC User Service

A simple gRPC service to manage user details. This service allows you to get user details, get multiple user details, and search for users based on criteria such as city, phone number, and marital status.
## Prerequisites

Make sure you have the following installed on your machine:

- [Go](https://golang.org/doc/install) (version 1.16+)
- [Protocol Buffers](https://grpc.io/docs/protoc-installation/) (protoc)
- [Docker](https://docs.docker.com/get-docker/) (for containerizing the application)

## Build and Run the Application

### Clone the Repository

git clone https://github.com/smitabharat/grpc-user-service.git
cd grpc-user-service

#Steps:
#1- Make sure you have protoc and the Go gRPC plugin installed. Generate the Go code from the .proto files:
   protoc --go_out=. --go-grpc_out=. proto/user_service.proto
#2- You can run the server locally using Go:
   go run main.go
#3 - Build the Docker image:
   docker build -t grpc-user-service .
#4- Run the Docker container:
   docker run -d -p 50051:50051 --name grpc-user-service grpc-user-service
#5- You can run the client to test the service:
   go run client/client.go
#6- The service provides the following endpoints:
   GetUserDetails :-
       Method: GetUserDetails
       Request: UserIdRequest
       Response: UserResponse
   SearchUsers
       Method: SearchUsers
       Request: SearchRequest
       Response: UsersResponse
#7 - Configuration Details
Port: The server listens on port 50051 by default. You can change this in the main.go file.
Users Data: The initial user data is hardcoded in the server/user_service.go file. You can modify this data as needed.


 
