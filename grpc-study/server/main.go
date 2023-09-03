package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-study/server/protos"
	"net"
)

type userService struct {
	protos.UnimplementedUserServiceServer
}

func (u *userService) Search(ctx context.Context, usr *protos.UserSearchRequest) (*protos.UserSearchResponse, error) {
	username := usr.Username
	return &protos.UserSearchResponse{
		Code:    1,
		Message: "success",
		UserInfo: &protos.UserInfo{
			Username: username,
			Password: "xxx",
			Type:     protos.UserType_USER_TYPE_ADMIN,
		},
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":9091")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	protos.RegisterUserServiceServer(grpcServer, &userService{})
	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}
