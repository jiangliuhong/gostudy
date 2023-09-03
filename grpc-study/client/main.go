package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-study/server/protos"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9091", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := protos.NewUserServiceClient(conn)
	request := protos.UserSearchRequest{Username: "test"}
	search, err := client.Search(context.Background(), &request)
	if err != nil {
		panic(err)
	}
	marshal, err := json.Marshal(search)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))
}
