package json_http

const SERVER_NAME = "HelloServer"

type HelloInterface interface {
	Hello(request string, response *string) error
}
