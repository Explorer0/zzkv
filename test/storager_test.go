package test_test

import (
	"fmt"
	"github.com/zzkv"
	"math/rand"
	"testing"
)

func TestPstStorager(t *testing.T) {
	s1 := zzkv.NewDefaultPstStorager()
	key := "fucker"
	values := []string{"12345879&……%%我要怎么说--+++!@#$%", "45879&……%%我要怎么说-", "fucker说什么", "bitcher zzkv渣渣键值对"}

	setErr := s1.Storage(key, []byte(values[0]))
	if setErr != nil {
		t.Fatal(fmt.Sprintf("failed to set. errMsg[%s]", setErr))
	}

	fetchVal := s1.Read(key)
	if string(fetchVal) != values[0] {
		t.Fatal("Inconsistent access data.")
	}

	setErr = s1.Storage(key, []byte(values[1]))
	if setErr != nil {
		t.Fatal(fmt.Sprintf("failed to set. errMsg[%s]", setErr))
	}

	fetchVal = s1.Read(key)
	if string(fetchVal) != values[1] {
		t.Fatal("Inconsistent access data.")
	}

	t.Log("---------------Test PersistentStorager PASS------------------")

}

func BenchmarkPstStorager(b *testing.B) {
	s1 := zzkv.NewDefaultPstStorager()
	key := "fucker"
	values := []string{"通过开源协作创建了大量优质编程教程", "已帮助全球数百万人学习编程", "非营利组织 freeCodeCamp", "成为开发者。在这里分享你的文章吧"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		setErr := s1.Storage(key, []byte(values[i%3]))
		if setErr != nil {
			b.Fatal(fmt.Sprintf("failed to set. errMsg[%s]", setErr))
		}

		fetchVal := s1.Read(key)
		if string(fetchVal) != values[i%3] {
			b.Fatal("Inconsistent access data.")
		}
	}
}

func BenchmarkPstStoragerThreadSafety(b *testing.B) {
	s1 := zzkv.NewDefaultPstStorager()
	key := "fucker"
	values := []string{"12345879&……%%我要怎么说--+++!@#$%", "45879&……%%我要怎么说-", "fucker说什么", "bitcher zzkv渣渣键值对"}
	valMap := make(map[string]bool)
	for _, val := range values {
		valMap[val] = true
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			idx := rand.Intn(4)
			setErr := s1.Storage(key, []byte(values[idx%4]))
			if setErr != nil {
				b.Fatal(fmt.Sprintf("failed to set. errMsg[%s]", setErr))
			}

			fetchVal := s1.Read(key)
			if _, ok := valMap[string(fetchVal)]; !ok && string(fetchVal) != "" {
				b.Fatal(fmt.Sprintf("Inconsistent access data. fetchedData:[%s]", string(fetchVal)))
			}
		}
	})
}


func TestCacheStrager(t *testing.T)  {
	s1 := zzkv.NewDefaultCacheStorager()
	key := "bitcher"
	values := []string{"12345879&……%%我要怎么说--+++!@#$%", "45879&……%%我要怎么说-", "fucker说什么", "bitcher zzkv渣渣键值对"}

	setErr := s1.Set(key, []byte(values[0]))
	if setErr != nil {
		t.Fatal(fmt.Sprintf("failed to set. errMsg[%s]", setErr))
	}

	fetchVal := s1.Get(key)
	if string(fetchVal) != values[0] {
		t.Fatal("Inconsistent access data.")
	}

	setErr = s1.Set(key, []byte(values[1]))
	if setErr != nil {
		t.Fatal(fmt.Sprintf("failed to set. errMsg[%s]", setErr))
	}

	fetchVal = s1.Get(key)
	if string(fetchVal) != values[1] {
		t.Fatal("Inconsistent access data.")
	}

	t.Log("---------------Test CacheStorager PASS------------------")

}

func BenchmarkCacheStorager (b *testing.B) {
	s1 := zzkv.NewDefaultCacheStorager()
	key := "bitcher"
	values := []string{"通过开源协作创建了大量优质编程教程", "已帮助全球数百万人学习编程", "非营利组织 freeCodeCamp", "成为开发者。在这里分享你的文章吧"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		setErr := s1.Set(key, []byte(values[i%4]))
		if setErr != nil {
			b.Fatal(fmt.Sprintf("failed to set. errMsg[%s]", setErr))
		}

		fetchVal := s1.Get(key)
		if string(fetchVal) != values[i%4] {
			b.Fatal("Inconsistent access data.")
		}
	}
}

func BenchmarkCacheStoragerThreadSafety(b *testing.B) {
	s1 := zzkv.NewDefaultCacheStorager()
	key := "fucker"
	values := []string{"12345879&……%%我要怎么说--+++!@#$%", "45879&……%%我要怎么说-", "fucker说什么", "bitcher zzkv渣渣键值对"}
	valMap := make(map[string]bool)
	for _, val := range values {
		valMap[val] = true
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			idx := rand.Intn(4)
			setErr := s1.Set(key, []byte(values[idx%4]))
			if setErr != nil {
				b.Fatal(fmt.Sprintf("failed to set. errMsg[%s]", setErr))
			}

			fetchVal := s1.Get(key)
			if _, ok := valMap[string(fetchVal)]; !ok && string(fetchVal) != "" {
				b.Fatal(fmt.Sprintf("Inconsistent access data. fetchedData:[%s]", string(fetchVal)))
			}
		}
	})
}


func TestStorager(t *testing.T) {
	s1 := zzkv.NewDefaultStorager()
	key := "fucker"
	values := []string{"12345879&……%%我要怎么说--+++!@#$%", "45879&……%%我要怎么说-", "fucker说什么", "bitcher zzkv渣渣键值对"}


	setErr := s1.Set(key, []byte(values[0]), true)
	if setErr != nil {
		t.Fatal(fmt.Sprintf("failed to set. errMsg[%s]", setErr))
	}

	fetchVal := s1.Get(key)
	if string(fetchVal) != values[0] {
		t.Fatal(fmt.Sprintf("Inconsistent access data. index:0, fetch value:%s", string(fetchVal)))
	}

	setErr = s1.Set(key, []byte(values[1]), false)
	if setErr != nil {
		t.Fatal(fmt.Sprintf("failed to set. errMsg[%s]", setErr))
	}

	fetchVal = s1.Get(key)
	if string(fetchVal) != values[1] {
		t.Fatal(fmt.Sprintf("Inconsistent access data. index:1, fetch value:%s", string(fetchVal)))
	}

	t.Log("---------------Test Storager PASS------------------")
}

func BenchmarkStoragerThreadSafety(b *testing.B) {
	s1 := zzkv.NewDefaultStorager()
	key := "bitcher"
	values := []string{"12345879&……%%我要怎么说--+++!@#$%", "45879&……%%我要怎么说-", "fucker说什么", "bitcher zzkv渣渣键值对"}
	valMap := make(map[string]bool)
	for _, val := range values {
		valMap[val] = true
	}
	
	b.RunParallel(func(pb *testing.PB) {
		idx := rand.Intn(4)

		setErr := s1.Set(key, []byte(values[idx%4]), true)
		if setErr != nil {
			b.Fatal(fmt.Sprintf("failed to set. errMsg[%s]", setErr))
		}

		fetchVal := s1.Get(key)
		if _, ok := valMap[string(fetchVal)]; !ok && string(fetchVal) != "" {
			b.Fatal(fmt.Sprintf("Inconsistent access data. fetchedData:[%s]", string(fetchVal)))
		}
	})
}



