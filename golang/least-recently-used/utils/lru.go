package utils

// cache holds up to k blocks (cache size)
// each block is a certain number of bytes (p.440)
// the input is a sequence of n memory requests
// that is, memory blocks b_1,b_2,...,b_n.
// we assume that b_i may (or may not) be equal 
// to b_j for i!=j
// If b_i is requested and is present on cache, we 
// have a cache hit. Otherwise, it is a cache miss.

import (
	"container/list"
)

// LRUCache represents a simple LRU cache
type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type entry struct {
	key   int
	value int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (c *LRUCache) Get(key int) bool {
	if el, found := c.cache[key]; found {
		c.list.MoveToFront(el)
		return true
	}
	return false
}

func (c *LRUCache) Put(key, value int) {
	if el, found := c.cache[key]; found {
		c.list.MoveToFront(el)
		el.Value = entry{key, value}
	} else {
		if c.list.Len() == c.capacity {
			back := c.list.Back()
			if back != nil {
				c.list.Remove(back)
				delete(c.cache, back.Value.(entry).key)
			}
		}
		el := c.list.PushFront(entry{key, value})
		c.cache[key] = el
	}
}