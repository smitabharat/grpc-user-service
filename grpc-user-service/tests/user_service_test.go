package tests

import (
	"context"
	"testing"

	"grpc-user-service/proto"
	"grpc-user-service/server"
)

func TestGetUserDetails(t *testing.T) {
	s := server.UserServiceServer{}
	req := &proto.UserIdRequest{Id: 1}
	res, err := s.GetUserDetails(context.Background(), req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if res.User.Fname != "Steve" {
		t.Errorf("Expected Steve, got %v", res.User.Fname)
	}
}

func TestGetUsersDetails(t *testing.T) {
	s := server.UserServiceServer{}
	req := &proto.UserIdsRequest{Ids: []int32{1, 2}}
	res, err := s.GetUsersDetails(context.Background(), req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(res.Users) != 2 {
		t.Errorf("Expected 2 users, got %v", len(res.Users))
	}
}

func TestSearchUsers(t *testing.T) {
	s := server.UserServiceServer{}
	req := &proto.SearchRequest{City: "LA"}
	res, err := s.SearchUsers(context.Background(), req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(res.Users) != 1 {
		t.Errorf("Expected 1 user, got %v", len(res.Users))
	}
	if res.Users[0].Fname != "Steve" {
		t.Errorf("Expected Steve, got %v", res.Users[0].Fname)
	}
}
 
