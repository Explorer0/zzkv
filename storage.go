package zzkv

import (
	"bytes"
	"io"
	"sync"
)

// 持久化存储
type PersistenceStorager interface {
	Storage(key string, reader io.Reader) error
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
		storageErr := s.outStorage.Storage(key, bytes.NewReader(val))
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

func (s DefaultPstStorager) Storage(key string, reader io.Reader) error {
	panic("implement me")
}

func (s DefaultPstStorager) Read(key string) []byte {
	panic("implement me")
}

func (s DefaultPstStorager) Delete(string) error {
	panic("implement me")
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


