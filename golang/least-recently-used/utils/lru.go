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

///// OFFLINE ALGORITHM (BELADY) //////
/*
LINKS:
	https://www.cs.cornell.edu/courses/cs4410/2015su/lectures/lec15-replacement.html
	https://www2.cs.uregina.ca/~hamilton/courses/330/notes/memory/page_replacement.html
	https://www.reddit.com/r/algorithms/comments/igkc1d/optimal_offline_caching_farthestinfuture_algorithm/
	https://courses.cs.washington.edu/courses/cse421/18au/lecture/lecture-8.pdf

*/


type LRUCacheOPT struct {
	capacity int 
	cache map[string]*list.Element
	list *list.List
	requests []string
	lastRequestNum int
}

func NewOPTCache(capacity int, requests []string) *LRUCacheOPT {
	return &LRUCacheOPT{
		capacity: capacity,
		cache: make(map[string]*list.Element),
		list: list.New(),
		requests: requests,
		lastRequestNum: -1, // no request made		
	}
}

func (c *LRUCacheOPT) Get(key string) (float32, bool) {
	// New request just made. Update counter!
	c.lastRequestNum++
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

func (c *LRUCacheOPT) CheckLateUse() *list.Element {
	
	// Variables
	elToEvict := c.list.Front()
	furthest := -1
	key := ""
	var diff int
	
	// Elements to iterate
	it := c.lastRequestNum + 1
	current := c.list.Front()
	
	for current != nil {
		// This cache element's key
		key = current.Value.(entryV1).key

		for it < len(c.requests){
			// If key appear ....
			if c.requests[it] == key {
				diff = it - c.lastRequestNum 
				// ... furthest in future ...
				if diff > furthest {
					// Update the element to be
					//  evicted and the furthest seen
					elToEvict = current
					furthest = diff
				}
			}
			it++
		}
		current = current.Next()
	}
	return elToEvict
}

func (c *LRUCacheOPT) Put(key string, value float32) {
	// Does the element exist?
	el, found := c.cache[key]
	if !found { // Does not
		// Check if cache is full
		if c.list.Len() == c.capacity { // It is. 
			// Who to evict? The item to be 
			// used farthest into the future.
			// The current request is c.lastRequestNum
			// We need to check which element from cache
			// will be used the farthest into future from
			// c.lastRequestNum on
			elToEvict := c.CheckLateUse()
			c.list.Remove(elToEvict)
			delete(c.cache, elToEvict.Value.(entryV1).key)
		}
		// Add to front list (recently used)
		el = c.list.PushFront(entryV1{key,value})
		// Save on hash
		c.cache[key] = el		
	} else {
		// Element does exist, move to front
		c.list.MoveToFront(el)
		el.Value = entryV1{key, value}
	}
	// and update the value
}