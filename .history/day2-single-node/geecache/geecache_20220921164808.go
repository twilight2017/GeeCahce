package geecache

import (
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
	getter    Getter
	mainCache cache
}

var (
	mu     sync.Mutex
	groups = make(map[string]*Group)
)

//NewGroup create a new instance of Group
func NewGroup(name string, cachBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter")
	}
}
