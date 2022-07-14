package gob

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type test struct {
	F1 string
	F2 int
}

func TestGobEncode(t *testing.T) {
	should := assert.New(t)
	gobBytes, err := GobEncode(&test{
		F1: "F1",
		F2: 10,
	})
	if should.NoError(err) {
		fmt.Println(gobBytes)
	}
	obj := test{}
	err = GobDecode(gobBytes, &obj)
	if should.NoError(err) {
		fmt.Println(obj)
	}
}

func TestGobDecode(t *testing.T) {

}
