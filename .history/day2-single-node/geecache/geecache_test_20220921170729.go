package geecache

import (
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
	gee := NewGroup("scores", 2<<10, GetterFunc(func(key string) {[]byte, error}{
		log.Println("[SlowDB] search key", key)
	}))
}