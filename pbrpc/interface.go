package pbrpc

const SERVER_NAME = "HelloServer"

type HelloInterface interface {
	Hello(request *Request, response *Response) error
}
