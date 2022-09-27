package geecache

import (
	"sync"
)

type cache struct {
	mu  sync.Mutex
	lru *lru.Ca
}
