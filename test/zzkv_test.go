package test

import (
	"fmt"
	"github.com/zzkv"
	"testing"
	"time"
)

type TestStt struct {
	X string	`json:"x"`
	Y string	`json:"y"`
}

func TestZzkv(t *testing.T)  {
	z1 := zzkv.NewDefault()
	t1 := TestStt{X:"fucker", Y:"shiter"}
	t2 := &TestStt{}
	key := "nba"

	err := z1.Set(key, t1, false)
	if err != nil {
		t.Fatal(fmt.Sprintf("Failed to set kv. errMsg[%s]", err))
	}

	err = z1.Get(key, t2)
	if err != nil {
		t.Fatal(fmt.Sprintf("Failed to get kv. errMsg[%s]", err))
	}

	if t1.X != t2.X || t1.Y != t2.Y {
		t.Fatal(fmt.Sprintf("Failed to test kv. Inconsistent access data."))
	}

	t.Log("----------------Test Zzkv PASS--------------------")
}

func TestZzkvClear(t *testing.T) {
	z1 := zzkv.NewDefault()
	t1 := TestStt{X:"fucker", Y:"shiter"}
	t2 := &TestStt{}
	key := "nba"

	err := z1.SetWithTTL(key, t1, true, 5)
	if err != nil {
		t.Fatal(fmt.Sprintf("Failed to set kv. errMsg[%s]", err))
	}
	time.Sleep(time.Second*10)
	err = z1.Get(key, t2)
	if err != nil {
		t.Fatal(fmt.Sprintf("Failed to get kv. errMsg[%s]", err))
	}

	t.Log("----------------Test ZzkvClear PASS--------------------")
}
