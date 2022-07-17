package main

import (
	"context"
	"fmt"
	"rpc-g7/grpc/simple/pb"
	"time"

	"google.golang.org/grpc"
)

func main() {
	// 1、建立网络链接
	conn, err := grpc.Dial(":1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)

	// 通过接口定义的方法就可以调用服务端对应的gRPC服务提供的方法
	req := &pb.Request{Value: "bob"}
	reply, err := client.Hello(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply.GetValue())

	// grpc会自动生成客户端的sdk
	client = pb.NewHelloServiceClient(conn)
	stream, err := client.Channel(context.Background())

	go func() {
		for {
			if err := stream.Send(&pb.Request{Value: "Alice"}); err != nil {
				panic(err)
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		response, err := stream.Recv()
		if err != nil {
			panic(err)
		}
		fmt.Println(response.Value)
	}
}
