package main

import (
	"fmt"
	"net/rpc"
	"rpc-g7/rpc_interface"
)

var _ rpc_interface.HelloService = (*HelloServiceClient)(nil)

func NewHelloServiceClient(network, address string) (*HelloServiceClient, error) {
	client, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}

	return &HelloServiceClient{
		client: *client,
	}, nil
}

type HelloServiceClient struct {
	client rpc.Client
}

func (h *HelloServiceClient) Hello(request string, response *string) error {
	// 1、建立socket连接
	return h.client.Call(fmt.Sprintf("%s.Hello", rpc_interface.SERVICE_NAME), request, &response)
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
