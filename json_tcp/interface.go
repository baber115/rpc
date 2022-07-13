package json_tcp

const (
	SERVICE_NAME = "HelloService"
)

type HelloService interface {
	Hello(request string, response *string) error
}
