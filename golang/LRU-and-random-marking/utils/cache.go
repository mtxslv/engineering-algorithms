package utils


import (
	"fmt"
	"errors"
	"container/list"
)

////////////////// LEAST RECENTLY USED /////////////////////

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


////////////////// RANDOM MARKING /////////////////////

var ErrCacheSizeTooLarge = errors.New("Cache size must be less than or equal to 16")


type RandomMarkingCache struct {
	capacity int
	markingByte uint16
	cache map[string]float32
}

func NewRandomMarkingCache(capacity int) (*RandomMarkingCache, error) {
	if capacity > 16 {
		return nil, fmt.Errorf("FAILED GIVEN:\n%w\n",ErrCacheSizeTooLarge)	
	}
	cachePtr := &RandomMarkingCache{
		capacity: capacity,
		markingByte: uint16(0),
		cache: make(map[string]float32),
	}
	return cachePtr, nil
}

func (c *RandomMarkingCache) Mark(i int) {
	c.markingByte = c.markingByte + 1 << i
}