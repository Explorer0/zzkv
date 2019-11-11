package test

import (
	"fmt"
	"github.com/zzkv"
	"testing"
)

func TestCompression(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Fatal(err)
		}
	}()

	data := "12345879&……%%我要怎么说--+++!@#$%"
	zipComp := zzkv.NewDefaultCompression()
	compressData := string(zipComp.Compress([]byte(data)))
	decompressData := string(zipComp.Decompress([]byte(compressData)))
	t.Log("------------Test Compression PASS------------")
	t.Log(fmt.Sprintf("compressed: %s", compressData))
	t.Log(fmt.Sprintf("decompressed: %s", decompressData))
}
