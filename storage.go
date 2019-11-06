package zzkv

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"sync"
)

const DefaultFileMode os.FileMode = 0666

// 持久化存储
type PersistentStorager interface {
	Storage(key string, value []byte) error
	Read(key string) []byte
	Delete(string) error
}

// 缓存
type CacheStorager interface {
	Set(string, []byte) error
	Get(string) []byte
	Erase(string) error
	IsExist(string) bool
}

// 存储器
type Storager struct {
	pstStorager		PersistentStorager
	cacheStorager 	CacheStorager
	storageMap 		map[string]bool
	sync.RWMutex
}

func (s *Storager) Set(key string, val []byte, sync bool) error {
	s.Lock()
	defer s.Unlock()

	// 硬件存储
	if sync {
		storageErr := s.pstStorager.Storage(key, val)
		if storageErr != nil {
			return storageErr
		}
		if _, ok := s.storageMap[key]; !ok {
			s.storageMap[key] = true
		}
	}

	//开启协程，单独写入缓存
	go func() {
		_ = s.cacheStorager.Set(key, val)
	}()

	return nil
}

func (s *Storager) Get(key string) []byte  {
	s.RLock()
	defer s.RUnlock()

	// 查看是否存在
	if _, ok := s.storageMap[key]; !ok {
		return nil
	}

	// 查看缓存是否命中
	if s.cacheStorager.IsExist(key) {
		return s.cacheStorager.Get(key)
	}

	// 缓存未命中，从持久化存储器取
	result := s.pstStorager.Read(key)

	// 开启协程，单独写缓存
	go func() {
		_ = s.cacheStorager.Set(key, result)
	}()

	return result
}



type DefaultPstStorager struct {
	sync.RWMutex
}

func (s *DefaultPstStorager) Storage(key string, value []byte) error {
	s.Lock()
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
	return nil
}

func (s *DefaultPstStorager) Read(key string) []byte {
	s.RLock()
	defer s.RUnlock()

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
	return result
}

func (s *DefaultPstStorager) Delete(key string) error {
	s.Lock()
	defer s.Unlock()
	fileName := fmt.Sprintf("%s.zzkv", key)
	cmd := exec.Command(fmt.Sprintf("rm %s", fileName))
	_, outErr := cmd.Output()
	return outErr
}



type DefaultCacheStorager struct {
	sync.RWMutex
}

func (s *DefaultCacheStorager) Set(string, []byte) error {
	panic("implement me")
}

func (s *DefaultCacheStorager) Get(string) []byte {
	panic("implement me")
}

func (s *DefaultCacheStorager) IsExist(string) bool {
	panic("implement me")
}

func (s *DefaultCacheStorager) Erase(string) error {
	panic("implement me")
}




func NewDefaultPstStorager() *DefaultPstStorager {
	return &DefaultPstStorager{
		sync.RWMutex{},
	}
}

func NewDefaultCacheStorager() *DefaultCacheStorager {
	return &DefaultCacheStorager{
		sync.RWMutex{},
	}
}


