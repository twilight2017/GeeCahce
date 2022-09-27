package geecache

import (
	"fmt"
	"log"
	"sync"
)

/*GeeCache不实现直接从数据源获取数据，原因有：
1.数据源的种类太多，无法一一实现
2.扩展性不好
如何从源头获取数据，应该是用户考虑完成实现的部分
*/

// A Getter loads data for a key
type Getter interface {
	//用[]byte存数据，是为了让它能支持任意数据格式
	Get(key string) ([]byte, error)
}

// A GetterFunc implements Getter with a function
//这个函数类型实现了接口方法，称之为接口类型的函数
type GetterFunc func(key string) ([]byte, error)

//Get implements Getter interface function
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

//A Group is a cache namespace and associated date loaded spread over
//Group是GeeCahe的核心数据结构，负责与用户的交互，并且控制缓存存值存储和获取的流程
type Group struct {
	name      string
	getter    Getter //缓存未命中时获取数据源的回调
	mainCache cache
}

var (
	mu     sync.RMMutex
	groups = make(map[string]*Group)
)

//NewGroup create a new instance of Group
func NewGroup(name string, cachBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{cacheBytes: cachBytes},
	}
	groups[name] = g
	return g
}

//GetGroup returns the named group previously created with NewGroup, or
// nil if there is no such group
func GetGroup(name string) *Group {
	mu.RLock()
	g := groups[name]
	mu.RUnlock()
	return g
}

//get value for a key from a cache
func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}

	//首先从maincache中查找缓存，如果存在则返回缓存值

	if v, ok := g.mainCache.get(key); ok {
		log.Println("[GeeCache hit]")
		return v, nil
	}

	//缓存不存在，调用load方法，load调用getLocally， getLocally调用用户回调函数g.getter.Get()获取源数据
	//并且将源数据添加到缓存mainCache中（通过populateCache方法）
	return g.load(key)
}

func (g *Group) load(key string) (value ByteView, err error) {
	return g.getLocally(key)
}

func (g *Group) getLocally(key string) (ByteView, error) {
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, nil
	}
	value := ByteView{b: cloneBytes(bytes)}
	g.populateCache(key, value)
	return value, nil
}

func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}
