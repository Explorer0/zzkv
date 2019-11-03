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

// 持久化存储
type PersistenceStorager interface {
	Storage(key string, value []byte) error
	Read(key string) []byte
	Delete(string) error
	RLock()
	RUnlock()
	Lock()
	Unlock()
}

// 缓存
type CacheStorager interface {
	Set(string, []byte) error
	Get(string) []byte
	IsExist(string) bool
	Erase(string) error
	RLock()
	RUnlock()
	Lock()
	Unlock()
}

// 抽象存储器
type Storager struct {
	outStorage		PersistenceStorager
	insideStorage 	CacheStorager
	storageMap 		map[string]bool
	sync.RWMutex
}

func (s *Storager) Set(key string, val []byte, sync bool) error {
	s.Lock()
	defer s.Unlock()

	// 缓存
	setErr := s.insideStorage.Set(key, val)
	if setErr != nil {
		return setErr
	}

	// 硬件存储
	if sync {
		storageErr := s.outStorage.Storage(key, val)
		if storageErr != nil {
			return storageErr
		}
		if _, ok := s.storageMap[key]; !ok {
			s.storageMap[key] = true
		}
	}

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
	if s.insideStorage.IsExist(key) {
		return s.insideStorage.Get(key)
	}

	// 缓存未命中，从持久化存储器取
	result := s.outStorage.Read(key)

	// 开启协程，单独写缓存
	go func() {
		s.insideStorage.Lock()
		defer s.insideStorage.Unlock()
		_ = s.insideStorage.Set(key, result)
	}()

	return result
}



type DefaultPstStorager struct {
	sync.RWMutex
}

func (s DefaultPstStorager) Storage(key string, value []byte) error {
	fileName := fmt.Sprintf("%s.zzkv", key)
	// 打开目标文件，不存在则创建
	fileHandle, openErr := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0444)
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

func (s DefaultPstStorager) Read(key string) []byte {
	fileName := fmt.Sprintf("%s.zzkv", key)
	fileHandle, openErr := os.OpenFile(fileName, os.O_RDONLY, 0444)
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

func (s DefaultPstStorager) Delete(key string) error {
	fileName := fmt.Sprintf("%s.zzkv", key)
	cmd := exec.Command(fmt.Sprintf("rm %s", fileName))
	_, outErr := cmd.Output()
	return outErr
}

func (s DefaultPstStorager) RLock() {
	s.RLock()
}

func (s DefaultPstStorager) RUnlock() {
	s.RUnlock()
}

func (s DefaultPstStorager) Lock() {
	s.Lock()
}

func (s DefaultPstStorager) Unlock() {
	s.Unlock()
}


type DefaultCacheStorager struct {
	sync.RWMutex
}

func (s DefaultCacheStorager) Set(string, []byte) error {
	panic("implement me")
}

func (s DefaultCacheStorager) Get(string) []byte {
	panic("implement me")
}

func (s DefaultCacheStorager) IsExist(string) bool {
	panic("implement me")
}

func (s DefaultCacheStorager) Erase(string) error {
	panic("implement me")
}

func (s DefaultCacheStorager) RLock() {
	s.RLock()
}

func (s DefaultCacheStorager) RUnlock() {
	s.RUnlock()
}

func (s DefaultCacheStorager) Lock() {
	s.Lock()
}

func (s DefaultCacheStorager) Unlock() {
	s.Unlock()
}




func NewDefaultPstStorager() DefaultPstStorager {
	return DefaultPstStorager{}
}

func NewDefaultCacheStorager() DefaultCacheStorager {
	return DefaultCacheStorager{}
}


