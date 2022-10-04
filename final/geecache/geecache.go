package geecache

import (
	"fmt"
	"sync"
)

// A Getter loads data for a key
type Getter interface {
	Get(key string) ([]byte, error)
}

//定义函数类型GetterFunc，并实现Getter接口的Get方法
//将这个函数称之为接口类型的函数
// A GetterFunc implements Getter with a function
type GetterFunc func(key string) ([]byte, error)

//Get implements Getter interface function
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

// A Group is a cache namespace and associated data loaded spred over
//一个Group是一个缓存的命名空间，name是其唯一标识符
type Group struct {
	name      string
	getter    Getter //缓存未命中时获取源数据的回调函数
	mainCache cache //一开始实现的并发缓存
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

//NewGroup create a new instance of Group
//实例化一个Group，并将其存在全局变量groups中
func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name: name,
		getter: getter,
		mainCache: cache{cacheBytes: cacheBytes}
	}
	groups[name] = g
	return g
}

//GetGroup returns the named group previouslv created with NewGroup, or
//nil if there's no such group
//获取特定名称的Group
func GetGroup(name string) *Group{
	mu.RLock()
	g := groups[name]
	mu.RUnlock()
	return g
}

//Get value for a key from cache
func(gg *Group) Get(key string) (Byteview, error){
	if key ==""{
		return ByteView{}, fmt.Errorf("key is required")
	}
	if v, ok := g.mainCache[key];ok{
		log.Println("[GeeCache hit]")
		return v, nil
	}
	return g.load(key)
}

func (g *Group) load(key string) (value ByteView, err error){
	return g.getLocally(key)
}

//从数据源去获取数据
func (g *Group)getLocally(key string) (ByteView, error){
	bytes, err := g.getter.Get(key)
	if err != nil{
		return ByteView{}, err
	}
	value := ByteView{b: cloneBytes(bytes)}
	g.populateCache(key, value)
	return value, nil
}

func (g *Group) populateCache(key string, value ByteView){
	g.mainCache.add(key, value)
}