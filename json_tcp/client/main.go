package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpc-g7/json_tcp"
)

var _ json_tcp.HelloService = (*HelloServiceClient)(nil)

func NewHelloServiceClient(network, address string) (*HelloServiceClient, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	// 1.建立socket连接
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return &HelloServiceClient{
		client: client,
	}, nil
}

type HelloServiceClient struct {
	client *rpc.Client
}

func (h *HelloServiceClient) Hello(request string, response *string) error {
	// 1、建立socket连接
	return h.client.Call(fmt.Sprintf("%s.Hello", json_tcp.SERVICE_NAME), request, &response)
}

func main() {
	client, err := NewHelloServiceClient("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	var response string
	err = client.Hello("alice", &response)
	if err != nil {
		panic(err)
	}
	// 打印调用完成的结果
	fmt.Println(response)
}
