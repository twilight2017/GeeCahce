package geecache

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestGetter(t *testing.T) {
	//借助GetterFunc的类型转换，将一个匿名的回调函数转换为了Getter类型的接口
	var f Getter = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})

	expect := []byte("key")
	//调用该接口的方法f.Get(ket string)，实际上就是在调用匿名回调函数
	if v, _ := f.Get("key"); !reflect.DeepEqual(v, expect) {
		t.Errorf("callback failed")
	}
}

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}


func TestGet(t *testing.T){
	loadCounts := make(map[string]int, len(db))
	gee := NewGroup("scores", 2<<10, GetterFunc(func(key string) ([]byte, error){
		log.Println("[SlowDB] search key", key)
		if v, ok := db[key];ok{
			if _, ok := loadCounts[key]; !ok{
				loadCounts[key] = 0
			}
			loadCounts[key] += 1
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s not exist", key)
	}))

	for k, v := range db{
		if view, err := gee.Get(k); err != nil || view.string() != v{
			t.Fatal("failed to get value of Tom")
		}
		if _, err := gee.Get(key); err != nil || loadCounts[k] >1{
			t.Fatal("cache %s miss", k )
		}
	}

	if view, err := gee.Get("unknown"; err == nil{
		t.Fatalf("the value of unknown should be emoty, but %s got", view)
	})
}