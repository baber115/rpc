package main

import (
	"fmt"
	"net"
	"net/rpc"
	"rpc-g7/json_tcp"
	"rpc-g7/pbrpc"
	"rpc-g7/pbrpc/codec/client"
)

var _ pbrpc.HelloInterface = (*HelloClient)(nil)

func NewHelloClient(network, address string) (*HelloClient, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	// 1.建立socket连接
	c := rpc.NewClientWithCodec(client.NewClientCodec(conn))
	return &HelloClient{
		client: c,
	}, nil
}

type HelloClient struct {
	client *rpc.Client
}

func (h *HelloClient) Hello(request *pbrpc.Request, response *pbrpc.Response) error {
	// 1、建立socket连接
	return h.client.Call(fmt.Sprintf("%s.Hello", json_tcp.SERVICE_NAME), request, response)
}

func main() {
	c, err := NewHelloClient("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	var response pbrpc.Response
	err = c.Hello(&pbrpc.Request{Value: "bb"}, &response)
	if err != nil {
		panic(err)
	}
	// 打印调用完成的结果
	fmt.Println(response.Value)
}
