// client/client.go
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "grpc-user-service/proto"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	
	// Call GetUserDetails method
	userResponse, err := c.GetUserDetails(ctx, &pb.UserIdRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	log.Printf("User: %s, City: %s", userResponse.GetUser().GetFname(), userResponse.GetUser().GetCity())

	// Call GetUsersDetails method
	usersResponse, err := c.GetUsersDetails(ctx, &pb.UserIdsRequest{Ids: []int32{1, 2}})
	if err != nil {
		log.Fatalf("could not get users: %v", err)
	}
	for _, user := range usersResponse.GetUsers() {
		log.Printf("User: %s, City: %s", user.GetFname(), user.GetCity())
	}

	// Call SearchUsers method with search criteria
	searchResponse, err := c.SearchUsers(ctx, &pb.SearchRequest{
		City:    "LA",   // search by city
		Phone:   1234567890, // search by phone number (optional)
		Married: true,  // search by marital status
	})
	if err != nil {
		log.Fatalf("could not search users: %v", err)
	}
	// Print the search results
	for _, user := range searchResponse.GetUsers() {
		log.Printf("User: %s, City: %s, Phone: %d, Height: %.1f, Married: %t",
            user.GetFname(), user.GetCity(), user.GetPhone(), user.GetHeight(), user.GetMarried())
	}
}
