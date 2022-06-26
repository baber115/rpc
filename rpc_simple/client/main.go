package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 1、建立socket连接
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	var response string
	err = client.Call("HelloService.Hello", "alice", &response)
	if err != nil {
		panic(err)
	}

	// 打印调用完成的结果
	fmt.Println(response)
}
