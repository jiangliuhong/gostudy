package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-study/server/protos"
	"net"
)

// userService 实现grpc服务端结构体
type userService struct {
	protos.UnimplementedUserServiceServer
}

// Search 实现grpc定义的rpc方法
func (u *userService) Search(ctx context.Context, usr *protos.UserSearchRequest) (*protos.UserSearchResponse, error) {
	username := usr.Username
	// 模拟输出
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
	// 定义grpc服务监听的端口
	listen, err := net.Listen("tcp", ":9091")
	if err != nil {
		panic(err)
	}
	// 初始化一个grpc服务
	grpcServer := grpc.NewServer()
	// 将定义的userService服务注册到grpc服务中
	protos.RegisterUserServiceServer(grpcServer, &userService{})
	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}
