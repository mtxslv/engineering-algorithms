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

// LRUCacheV0 represents a simple LRU cache
type LRUCacheV0 struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type entryV0 struct {
	key   int
	value int
}

func NewLRUCacheV0(capacity int) *LRUCacheV0 {
	return &LRUCacheV0{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (c *LRUCacheV0) Get(key int) bool {
	// If the key exists in the hashmap,
	// move it to front (recently used)
	// and return true
	if el, found := c.cache[key]; found {
		c.list.MoveToFront(el)
		return true
	}
	return false
}

func (c *LRUCacheV0) Put(key, value int) {
	// The element exist, move to front
	// (recently used) and update the hash
	if el, found := c.cache[key]; found {
		c.list.MoveToFront(el)
		el.Value = entryV0{key, value}
	} else {
		if c.list.Len() == c.capacity {
			back := c.list.Back()
			if back != nil {
				c.list.Remove(back)
				delete(c.cache, back.Value.(entryV0).key)
			}
		}
		el := c.list.PushFront(entryV0{key, value})
		c.cache[key] = el
	}
}