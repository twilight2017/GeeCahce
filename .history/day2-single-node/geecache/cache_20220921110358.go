package geecache

import (
	"day2-single-node/geecache/lru"
	"sync"
)

type cache struct {
	mu  sync.Mutex
	lru *lru.Ca
}
