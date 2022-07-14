package pb_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"rpc-g7/protobuf/pb"
	"testing"
)

func TestMarshal(t *testing.T) {
	should := assert.New(t)
	str := &pb.String{
		Value: "Hello",
	}
	// object --> protobuf ]byte
	pbBytes, err := proto.Marshal(str)
	if should.NoError(err) {
		fmt.Println(pbBytes)
	}
	// protobuf []byte --> object
	obj := pb.String{}
	err = proto.Unmarshal(pbBytes, &obj)
	if should.NoError(err) {
		fmt.Println(obj)
	}
}
