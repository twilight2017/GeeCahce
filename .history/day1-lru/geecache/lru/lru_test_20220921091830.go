package lru

import (
	"fmt"
	"reflect"
	"testing"
)

type String string

func (d String) Len() int {
	return len(d)
}

func TestGet(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key1", String("1234"))
	if v, ok := lru.Get("key1"); !ok || string(v.(String)) != "1234" {
		t.Fatalf("cache hit key1=1234 failed")
	} else {
		fmt.Sprintf("success")
	}
	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	} else {
		fmt.Sprintf("key2 test success")
	}
}

func TestRemoveoldest(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "key3"
	v1, v2, v3 := "v1", "v2", "v3"
	cap := len(k1 + k2 + v1 + v2)
	lru := New(int64(cap), nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))
	if _, ok := lru.Get("key1"); ok || lru.Len() != 2 {
		t.Fatalf("Removeoldest key1 failed")
	}
}

//测试回调函数是否起作用
// func TestOnEvicated(t *testing.T) {
// 	//回调函数在某条记录被移除时调用
// 	keys := make([]string, 0)
// 	callback := func(key string, value Value) {
// 		keys = append(keys, key)
// 	}
// 	lru := New(int64(10), callback)
// 	lru.Add("key1", String("1"))
// 	lru.Add("key2", String("12"))
// 	lru.Add("key3", String("123"))
// 	lru.Add("key4", String("1234"))

// 	expect := []string{"key1", "key2"}
// 	if !reflect.DeepEqual(expect, keys) {
// 		t.Fatalf("Call OnEvicated failed, expect keys equal to %s", expect)
// 	}
// }
