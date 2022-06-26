package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"rpc-g7/rpc_interface"
)

var _ rpc_interface.HelloService = (*HelloService)(nil)

type HelloService struct {
}

// request 请求，response 响应
// request 输入name
// response 返回 hello, name
func (h *HelloService) Hello(request string, response *string) error {
	*response = fmt.Sprintf("Hello, %s", request)

	return nil
}

// main里面编写Server
func main() {
	// 把rpc对外暴露的对象注册到rpc框架内部
	rpc.RegisterName(rpc_interface.SERVICE_NAME, &HelloService{})

	// 1、准备socket
	// 建立一个唯一的TCP链接
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTcp error:", err)
	}

	// 2、获取链接
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		// 每个客户端单独启动一个routine来处理
		go rpc.ServeConn(conn)
	}
}
