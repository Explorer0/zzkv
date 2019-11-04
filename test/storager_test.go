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
	t.Log("PersistentStorager test---------------------------")

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

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			idx := rand.Intn(3)
			setErr := s1.Storage(key, []byte(values[idx%3]))
			if setErr != nil {
				b.Fatal(fmt.Sprintf("failed to set. errMsg[%s]", setErr))
			}
			_ = s1.Read(key)
		}
	})
}


