package geecache

import (
	"reflect"
	"testing"
)

func TestGetter(t *testing.T) {
	//借助GetterFunc的类型转换，将一个匿名的回调函数转换为了Getter类型的接口
	var f Getter = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})

	expect := []byte("key")
	if v, _ := f.Get("key"); !reflect.DeepEqual(v, expect) {
		t.Errorf("callback failed")
	}
}
