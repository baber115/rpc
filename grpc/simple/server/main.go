package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"rpc-g7/grpc/simple/middleware"
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

func (h *HelloServiceServer) Channel(Channel pb.HelloService_ChannelServer) error {
	// 循环接收客户端发送的请求
	for {
		request, err := Channel.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		err = Channel.Send(&pb.Response{
			Value: "Hello " + request.Value,
		})
		if err != nil {
			return err
		}
	}
}

func main() {
	// new grpc server

	// 添加认证中间件
	auth := grpc.UnaryInterceptor(middleware.GrpcAuthUnaryServerInterceptor())
	streamAuth := grpc.StreamInterceptor(middleware.GrpcAuthStreamServerInterceptor())
	server := grpc.NewServer(auth, streamAuth)
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
