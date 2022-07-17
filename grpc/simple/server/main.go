package main

import (
	"context"
	"fmt"
	"net"
	"rpc-g7/grpc/simple/pb"

	"google.golang.org/grpc"
)

type HelloServiceServer struct {
	pb.UnimplementedHelloServiceServer
}

func (h *HelloServiceServer) Hello(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	return &pb.Response{
		Value: fmt.Sprintf("Hello, %s", request.Value),
	}, nil
}

func main() {
	// new grpc server
	server := grpc.NewServer()

	//传参 grpc server 和 实现类
	pb.RegisterHelloServiceServer(server, &HelloServiceServer{})

	Listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	// 监听soce，http2内置
	if err := server.Serve(Listener); err != nil {
		panic(err)
	}
}
