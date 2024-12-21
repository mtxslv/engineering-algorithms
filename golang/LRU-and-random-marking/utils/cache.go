package utils

import (
	"container/list"
	"errors"
	"fmt"
	"math/rand"
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
var ErrAssignmentOutOfBounds = errors.New("Trying to reach a bit beyond cache capacity")

type RandomMarkingCache struct {
	capacity int
	markingByte uint16
	cache map[string]float32
	cacheOrder []string
}

func NewRandomMarkingCache(capacity int) (*RandomMarkingCache, error) {
	if capacity > 16 {
		return nil, fmt.Errorf("FAILED GIVEN:\n%w\n",ErrCacheSizeTooLarge)	
	}
	cachePtr := &RandomMarkingCache{
		capacity: capacity,
		markingByte: uint16(0),
		cache: make(map[string]float32),
		cacheOrder: make([]string,capacity),
	}
	return cachePtr, nil
}

/*i starts at 0*/
func (c *RandomMarkingCache) Mark(i int) error {

	if i >= c.capacity {
		return fmt.Errorf("Failed given:\n%w\n",ErrAssignmentOutOfBounds)
	}
	c.markingByte = c.markingByte | 1 << i
	return nil
}

/*i starts at 0*/
func (c *RandomMarkingCache) Unmark(i int) error {
	if i >= c.capacity {
		return fmt.Errorf("Failed given:\n%w\n",ErrAssignmentOutOfBounds)
	}	
	mask := 65534<<i  + 1<<i - 1 
	c.markingByte = c.markingByte & uint16(mask)
	return nil
} 

func (c *RandomMarkingCache) AllMarked() bool {
	if c.markingByte == 1<< uint16(c.capacity) -1 {
		return true
	} else {
		return false
	}
} 

func (c *RandomMarkingCache) KeyPosition(key string) int {
	for it, k := range c.cacheOrder {
		if k == key {
			return it
		}
	}
	return -1
}

// b is the key, i guess
func (c *RandomMarkingCache) RandomizedMarking(b string) {
	// is b block in the cache?
	_, found := c.cache[b]
	if found {
		// Where is b?
		it := c.KeyPosition(b) // CHECAR DEPOIS SE NÃO É -1
		c.Mark(it)
	} else {
		// are all blocks marked?
		if c.AllMarked() {
			// unmark all
			c.markingByte = 0
		}
		// Select an unmarked block uniformly at random
		toEvict := rand.Intn(c.capacity)
		// Evict block u
		keyToEvict := c.cacheOrder[toEvict]
		delete(c.cache,keyToEvict)
		// place block b into the cache
		c.cache[b] = 0.0 // FIX IT LATER
		c.cacheOrder[toEvict] = b
	}
}

func (c *RandomMarkingCache) Get(key string) (float32, bool) {
	// Does the element exist?
	el, found := c.cache[key]
	if !found {
		// No, return nothing
		return 0.0,false
	} else {
		// It is, let's mark it
		it := c.KeyPosition(key)  // CHECAR DEPOIS SE NÃO É -1
		c.Mark(it)
		return el, true
	}
}

// func (c *RandomMarkingCache) Put(key string, value float32) {
// 	// Does the element exist?
// 	el, found := c.cache[key]
// 	if found {
// 		// I guess I'll need to update value?
// 		c.Mark()

// 	}
// }


func (c *RandomMarkingCache) UnmarkedBitsArray() []int {
	var unmarkedPositions []int 
	it := 0
	for it < c.capacity {
		mask := 1 << it
		maskedBit := c.markingByte & uint16(mask)
		if maskedBit == 0 {
			unmarkedPositions = append(unmarkedPositions, it)
		}
		it++
	}
	return unmarkedPositions
}

func (c *RandomMarkingCache) SelectFromUnmarked() int {
	unmarkedPositions := c.UnmarkedBitsArray()
	if len(unmarkedPositions) == 0 {
		return -1
	} else {
		return unmarkedPositions[
			rand.Intn(len(unmarkedPositions)),
		]
	}
}
