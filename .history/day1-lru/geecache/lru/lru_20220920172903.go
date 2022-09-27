package lru

import "container/list"

//Cache is a LRU cache. It is not safe for concurrent access
type Cache struct {
	maxBytes int64                    //允许使用的最大内存
	nbytes   int64                    //当前已经使用的内存
	ll       *list.list               //Go语言标准库提供的双向链表
	cache    map[string]*list.Element //值是双向链表中节点的指针
	//optional and executed when an entry is purged
	OnEvicted func(key string, value Value) //某条记录被移除时的回调函数
}

//双向链表节点的数据类型
type entry struct {
	key   string
	value Value //值是实现了Value接口的任意类型，该接口只包含了一个方法Len()，返回值所占用的内存的大小
}

type Value interface {
	Len()
}

// New() is the constructor of Cache
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

//查找功能
//1.从字典中找到对应的双向链表的节点
//2.将该节点移至队尾
//Get look ups a key's value
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele) //键对应的链表节点存在，将对应节点移动到队尾，并返回查找到的值
		kv := ele.Value.(*entry)
		return kv.value, true
	}
}

//删除，移除最少访问的节点，即队首节点
//Remove the oldest item
func (c *Cache) RemoveOldest() {
	//取到队首节点
	ele := c.ll.back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len)
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}

}
