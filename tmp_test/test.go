package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"sync"
	"time"
)

const DefaultFileMode os.FileMode = 0666

type DefaultPstStorager struct {
	*sync.Mutex
}

func (s DefaultPstStorager) Storage(key string, value []byte) error {
	s.Lock()
	fmt.Printf("s:%p\n", &s)
	fmt.Println("locking...")
	defer s.Unlock()
	fileName := fmt.Sprintf("%s.zzkv", key)
	// 打开目标文件，不存在则创建
	fileHandle, openErr := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, DefaultFileMode)
	if openErr != nil {
		return openErr
	}

	// 将value写入文件
	_, writeErr := io.Copy(io.Writer(fileHandle), bytes.NewReader(value))
	if writeErr != nil {
		return writeErr
	}
	_ = fileHandle.Close()
	fmt.Println("unlocking...")
	return nil
}

func (s DefaultPstStorager) Read(key string) []byte {
	s.Lock()
	fmt.Printf("s:%p\n", &s)
	fmt.Println("Rlocking...")
	defer s.Unlock()

	fileName := fmt.Sprintf("%s.zzkv", key)
	fileHandle, openErr := os.OpenFile(fileName, os.O_RDONLY, DefaultFileMode)
	if openErr != nil {
		panic(fmt.Sprintf("Occur fatal error while opening file. errMsg[%s]", openErr))
	}

	//从文件中读取value
	result, readErr := ioutil.ReadAll(fileHandle)
	if readErr != nil {
		panic(fmt.Sprintf("Occur fatal error while read file. errMsg[%s]", readErr))
	}
	_ = fileHandle.Close()
	fmt.Println("Runlocking...")
	return result
}

func NewDefaultPstStorager() DefaultPstStorager {
	return DefaultPstStorager{
		&sync.Mutex{},
	}
}

func main()  {
	s2 := NewDefaultPstStorager()
	fmt.Printf("s2:%p\n", &s2)
	key := "xxx"
	values := []string{"12345879&……%%我要怎么说测试--+++!@#$%", "45879&……%%测试我要怎么说-", "fucker测试", "bitcher zzkv渣渣键值对"}
	valMap := make(map[string]bool)
	for _, val := range values {
		valMap[val] = true
	}

	for i := 0; i < 100; i++{
		go func() {
			idx := rand.Intn(4)
			err := s2.Storage(key, []byte(values[idx%4]))
			if err != nil {
				panic(err)
			}

			fetchVal := s2.Read(key)
			if _, ok := valMap[string(fetchVal)]; !ok && string(fetchVal) != "" {
				fmt.Printf("Inconsistent access data. fetchedData:[%s]\n", string(fetchVal))
			}
		}()
	}
	time.Sleep(time.Second*5)
}
