package zzkv

import (
	"github.com/gogf/gf/encoding/gcompress"
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


// 默认压缩器不做任何数据变动
type DefaultCompression struct {}

func (compress *DefaultCompression) Compress(originalVal []byte) []byte {
	data, err := gcompress.Gzip(originalVal)
	if err != nil {
		panic(err)
	}
	return data
}

func (compress *DefaultCompression) Decompress(compressedVal []byte) []byte {
	data, err := gcompress.UnGzip(compressedVal)
	if err != nil {
		panic(err)
	}

	return data
}

func NewDefaultCompression() *DefaultCompression {
	return &DefaultCompression{}
}