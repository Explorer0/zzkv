package zzkv

import (
	"sync"
	"time"
)

const IntervalDuration = 60		//60s


type Clear struct {
	ttlMap 		map[string]int64
	sync.Mutex
}

func NewDefaultClear() *Clear {
	return &Clear{
		make(map[string]int64),
		sync.Mutex{},
	}
}

// 标记过期时间
func (clear *Clear) Mark(key string, expireSeconds int64) {
	clear.Lock()
	defer clear.Unlock()
	clear.ttlMap[key] = expireSeconds
}

// 定时删除函数
func (clear *Clear) TimingErase(storager *Storager) {
	expiredKeyList := make([]string, 0)

	clear.Lock()
	// 寻找过期key，并更新过期时间
	for key, time := range clear.ttlMap {
		if time - IntervalDuration <= 0 {
			expiredKeyList = append(expiredKeyList, key)
		}
		clear.ttlMap[key] = time - IntervalDuration
	}
	// 删除过期key
	for _, key := range expiredKeyList {
		delete(clear.ttlMap, key)
	}
	clear.Unlock()

	// 删除所有过期kv
	for _, key := range expiredKeyList {
		storager.Erase(key)
	}

}

func (clear *Clear) Run(storager *Storager) {
	ticker := time.NewTicker(time.Second * IntervalDuration)
	go func() {
		for range ticker.C {
			clear.TimingErase(storager)
		}
	}()
}


