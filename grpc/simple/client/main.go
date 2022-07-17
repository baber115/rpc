package main

import (
	"context"
	"fmt"
	"rpc-g7/grpc/simple/pb"

	"google.golang.org/grpc"
)

func main() {
	// 1、建立网络链接
	conn, err := grpc.DialContext(context.Background(), ":1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	// grpc会自动生成客户端的sdk
	client := pb.NewHelloServiceClient(conn)
	response, err := client.Hello(context.Background(), &pb.Request{Value: "Alice"})
	if err != nil {
		panic(err)
	}

	fmt.Print(response.Value)
}
