package lru

import "container/list"

//Cache is a LRU cache. It is not safe for concurrent access
type Cache struct {
	maxBytes int64
	nbytes   int64
	ll       *list.list
	cache    map[string]*list.Element
	//optional and executed when an entry is purged
	OnEvicted func(key string, value Value)
}
