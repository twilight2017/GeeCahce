package cache

import (
	"geecache/lru"
	"sync"
)

type cache struct {
	mu  sync.Mutex
	lru *lru.Ca
}
