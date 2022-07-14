package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"rpc-g7/json_tcp"
	"rpc-g7/pbrpc"
	"rpc-g7/pbrpc/codec/server"
)

var _ pbrpc.HelloInterface = (*HelloService)(nil)

type HelloService struct {
}

// request 请求，response 响应
// request 输入name
// response 返回 hello, name
func (h *HelloService) Hello(request *pbrpc.Request, response *pbrpc.Response) error {
	response.Value = fmt.Sprintf("Hello, %s", request.Value)

	return nil
}

// main里面编写Server
func main() {
	// 把rpc对外暴露的对象注册到rpc框架内部
	rpc.RegisterName(json_tcp.SERVICE_NAME, &HelloService{})

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
		// server端采用json进行编解码，类似json.unmarshal和marshal
		go rpc.ServeCodec(server.NewServerCodec(conn))
	}
}
