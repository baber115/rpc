package main

import (
	"context"
	"fmt"
	"rpc-g7/grpc/simple/middleware"
	"rpc-g7/grpc/simple/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	case1()
	fmt.Println("case1 successfull")
	case2()
	fmt.Println("case2 successfull")
	case3()
	fmt.Println("case3 successfull")
}

/**
基础发出和接收
通过context 传入header参数（client_id,client_secret）
**/
func case1() {
	conn, err := grpc.DialContext(context.Background(), ":1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)
	ctx := metadata.NewOutgoingContext(context.Background(), NewClientCredential("admin", "123456"))
	// 通过接口定义的方法就可以调用服务端对应的gRPC服务提供的方法
	reply, err := client.Hello(
		ctx,
		&pb.Request{Value: "case1"},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply.GetValue())
}

// 用户端添加认证凭证
func NewClientCredential(ak, sk string) metadata.MD {
	return metadata.MD{
		middleware.ClientHeaderKey: []string{ak},
		middleware.ClientSecretKey: []string{sk},
	}
}

/**
WithPerRPCCredentials
**/
func case2() {
	conn, err := grpc.Dial(":1234", grpc.WithInsecure(), grpc.WithPerRPCCredentials(middleware.NewClientAuthentication("admin", "123456")))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)
	// 通过接口定义的方法就可以调用服务端对应的gRPC服务提供的方法
	reply, err := client.Hello(
		context.Background(),
		&pb.Request{Value: "case2"},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply.GetValue())
}

// 循环发出和循环接收
func case3() {
	conn, err := grpc.Dial(":1234", grpc.WithInsecure(), grpc.WithPerRPCCredentials(middleware.NewClientAuthentication("admin", "123456")))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// grpc会自动生成客户端的sdk
	client := pb.NewHelloServiceClient(conn)
	stream, err := client.Channel(context.Background())

	go func() {
		for {
			if err := stream.Send(&pb.Request{Value: "case3"}); err != nil {
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
