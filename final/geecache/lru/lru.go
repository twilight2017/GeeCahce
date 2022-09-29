package lru

import "container/list"

//Cache is a LRU cache. It is not safe for concurrent cache
type Cache struct {
	maxBytes int64                    //允许使用的最大内存
	nbytes   int64                    //当前已经使用的内存
	ll       *list.List               //Go语言标准库实现的双向链表
	cache    map[string]*list.Element //双向链表中对应节点的指针
	//optional and executed when an entry is purged
	OnEvicted func(key string, value Value) //某条记录被移除时的回调函数
}

//双向链表节点的数据类型，保存key是为了：
//删除队首节点时，需要用key从字典中删除对应的映射
type entry struct {
	key   string
	value Value
}

//Value use len to count how many bytes it takes
type Value interface {
	Len() int
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

//查找：从字典中找到对应的双向链表的节点，随后将该节点移到队尾
//Get look ups a key's value
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

//删除：移除队首节点即可
func (c *Cache) RemoveOldest() {
	//取到队首节点
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		//从字典中删除该节点的映射关系
		delete(c.cache, kv.key)
		//更新当前所用内存
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

//新增/修改
// Add adds a value to the cache
func (c *Cache) Add(key string, value Value) {
	//更新
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value //更新value域的值
	} else {
		//1.将新增节点推入双向链表
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest() //超过内存最大限制时删除队首元素
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
