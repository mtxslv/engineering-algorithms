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

/////////////// NEW LRU ///////////////

type entryV1 struct {
	key string 
	value float32
	// elementPointer *list.Element
}

type LRUCacheV1 struct {
	capacity int
	cache map[string]*list.Element
	list *list.List
}

func NewLRUCacheV1(capacity int) *LRUCacheV1 {
	return &LRUCacheV1{
		capacity: capacity,
		cache: make(map[string]*list.Element),
		list: list.New(),
	}
}

func (c *LRUCacheV1) Put(key string, value float32){
	// Does the element exist?
	el, found := c.cache[key]
	if !found {
		// It does not! Check if the cache is full
		if c.list.Len() == c.capacity {
			// Is full, so let's evict last item
			back := c.list.Back() // least used item
			if back != nil {
				c.list.Remove(back) // evict from usage list
				delete(c.cache,back.Value.(entryV1).key)
			}
		}
		// Add to front list (recently used)
		el = c.list.PushFront(entryV1{key,value})
		// Save on hash
		c.cache[key] = el
	}  else {
		// Element does exist, move to front
		c.list.MoveToFront(el)
		el.Value = entryV1{key, value}
	}
	// and update the value
}

func (c *LRUCacheV1) Get(key string) (float32, bool) {
	// Does the element exist?
	el, found := c.cache[key]	
	if !found{
		// No, return nothing
		return 0.0,false
	} else {
		// Since it is recently used
		// Move to front
		c.list.MoveToFront(el)
		value := el.Value.(entryV1).value
		return value, true
	}
}