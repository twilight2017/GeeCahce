package geecache

import (
	"lru"
	"sync"
)

type cache struct {
	mu  sync.Mutex
	lru *lru
}
