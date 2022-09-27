package cache

import (
	"sync"
)

type cache struct {
	mu  sync.Mutex
	lru *lru.Ca
}
