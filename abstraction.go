package zzkv

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var NoneError = errors.New("nil value")

// 压缩接口
type Compression interface {
	Compress([]byte) []byte
	Decompress([]byte) []byte
}

// 数据抽象层
func Serialize(originalVal interface{}) ([]byte, error){
	if originalVal == nil {
		return nil, NoneError
	}
	return json.Marshal(originalVal)
}

func Deserialize(serializedBytes []byte, val interface{}) error {
	if len(serializedBytes) <= 0 {
		return NoneError
	}

	return json.Unmarshal(serializedBytes, val)
}

type DefaultCompression struct {

}

func (DefaultCompression) Compress([]byte) []byte {
	panic("implement me")
}

func (DefaultCompression) Decompress([]byte) []byte {
	panic("implement me")
}

func NewDefaultCompression() DefaultCompression {
	return DefaultCompression{}
}