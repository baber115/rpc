package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpc-g7/json_http"
)

var _ json_http.HelloInterface = (*HelloServer)(nil)

type HelloServer struct {
}

// 方法一
func (h *HelloServer) Hello(request string, response *string) error {
	*response = fmt.Sprintf("Hello, %s", request)

	return nil
}

type CalcRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

// 方法二
func (h *HelloServer) Calc(request *CalcRequest, response *int) error {
	*response = request.A + request.B

	return nil
}

type RPCReadWriterCloser struct {
	io.Writer
	io.ReadCloser
}

func NewRPCReadWriterCloser(w http.ResponseWriter, r *http.Request) *RPCReadWriterCloser {
	return &RPCReadWriterCloser{
		w,
		r.Body,
	}
}

func main() {
	rpc.RegisterName(json_http.SERVER_NAME, &HelloServer{})
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		rpc.ServeCodec(jsonrpc.NewServerCodec(NewRPCReadWriterCloser(w, r)))
	})
	http.ListenAndServe(":1234", nil)
}
