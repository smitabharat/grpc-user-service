package server

import (
	"context"
	"errors"
	"sync"

	"grpc-user-service/models"
	"grpc-user-service/proto"
)

var users = []models.User{
	{ID: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
	{ID: 2, Fname: "John", City: "NY", Phone: 9876543210, Height: 6.1, Married: false},
}

type UserServiceServer struct {
	proto.UnimplementedUserServiceServer
	mutex sync.Mutex
}

func (s *UserServiceServer) GetUserDetails(ctx context.Context, req *proto.UserIdRequest) (*proto.UserResponse, error) {
	for _, user := range users {
		if user.ID == req.Id {
			return &proto.UserResponse{User: &proto.User{
				Id:      user.ID,
				Fname:   user.Fname,
				City:    user.City,
				Phone:   user.Phone,
				Height:  user.Height,
				Married: user.Married,
			}}, nil
		}
	}
	return nil, errors.New("user not found")
}

func (s *UserServiceServer) GetUsersDetails(ctx context.Context, req *proto.UserIdsRequest) (*proto.UsersResponse, error) {
	var responseUsers []*proto.User
	for _, id := range req.Ids {
		for _, user := range users {
			if user.ID == id {
				responseUsers = append(responseUsers, &proto.User{
					Id:      user.ID,
					Fname:   user.Fname,
					City:    user.City,
					Phone:   user.Phone,
					Height:  user.Height,
					Married: user.Married,
				})
			}
		}
	}
	return &proto.UsersResponse{Users: responseUsers}, nil
}

func (s *UserServiceServer) SearchUsers(ctx context.Context, req *proto.SearchRequest) (*proto.UsersResponse, error) {
	var responseUsers []*proto.User
	for _, user := range users {
		if (req.City == "" || user.City == req.City) &&
			(req.Phone == 0 || user.Phone == req.Phone) &&
			(req.Married == user.Married) {
			responseUsers = append(responseUsers, &proto.User{
				Id:      user.ID,
				Fname:   user.Fname,
				City:    user.City,
				Phone:   user.Phone,
				Height:  user.Height,
				Married: user.Married,
			})
		}
	}
	return &proto.UsersResponse{Users: responseUsers}, nil
}
