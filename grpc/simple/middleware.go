package simple

import (
	"github.com/infraboard/mcube/logger"
	"go.uber.org/zap"
)

const (
	ClientHeaderKey = "client-id"
	ClientSecretKey = "client-secret"
)

// internal todo
type grpcAuther struct {
	log &logger.Logger
}

func newGrpcAuther() *grpcAuther {
	return &grpcAuther{
		log: zap.L().Named("grpc auther"),
	}
}

// func (g grpcAuth) Auth(
// 	ctx context.Context,
// 	req interface{},
// 	info *grpc.UnaryServerInfo,
// 	handler grpc.UnaryHandler,
// ) (resp interface{}, err error) {

// }
