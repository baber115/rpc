package codec

import (
	"bytes"
	"encoding/gob"
)

func GobEncode(obj interface{}) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	// 编码后的结果输出到buf里面
	encode := gob.NewEncoder(buf)
	// encode obj对象
	if err := encode.Encode(obj); err != nil {
		return []byte{}, err
	}

	return buf.Bytes(), nil
}

func GobDecode(data []byte, obj interface{}) error {
	reader := bytes.NewReader(data)
	decode := gob.NewDecoder(reader)

	return decode.Decode(obj)
}
