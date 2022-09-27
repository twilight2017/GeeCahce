package cache

import (
	"lru"
	"sync"
)

type cache struct {
	mu  sync.Mutex
	lru *lru.Ca
}
