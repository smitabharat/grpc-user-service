package main

import (
	"log"
	"net"

	"grpc-user-service/proto"
	"grpc-user-service/server"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterUserServiceServer(s, &server.UserServiceServer{})

	log.Printf("Server is running at %v", listener.Addr().String())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
