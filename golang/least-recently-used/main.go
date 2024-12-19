// package main

// import "utils/utils"

// func main() {
// 	utils.SimulationWithRandomness()
// }

package main

import (
	"container/list"
	"fmt"
)

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

type OptimalCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
	requests []int
}

func NewOptimalCache(capacity int, requests []int) *OptimalCache {
	return &OptimalCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
		requests: requests,
	}
}

func (c *OptimalCache) Get(key int) bool {
	if _, found := c.cache[key]; found {
		return true
	}
	return false
}

func (c *OptimalCache) Put(key, value, currentIndex int) {
	if el, found := c.cache[key]; found {
		el.Value = entry{key, value}
	} else {
		if c.list.Len() == c.capacity {
			c.evict(currentIndex)
		}
		el := c.list.PushBack(entry{key, value})
		c.cache[key] = el
	}
}

func (c *OptimalCache) evict(currentIndex int) {
	var farthestIndex, evictKey int = -1, -1
	for el := c.list.Front(); el != nil; el = el.Next() {
		key := el.Value.(entry).key
		index := c.findNextUse(key, currentIndex)
		if index == -1 {
			farthestIndex, evictKey = index, key
			break
		} else if index > farthestIndex {
			farthestIndex, evictKey = index, key
		}
	}
	if evictKey != -1 {
		if el, found := c.cache[evictKey]; found {
			c.list.Remove(el)
			delete(c.cache, evictKey)
		}
	}
}

func (c *OptimalCache) findNextUse(key, currentIndex int) int {
	for i := currentIndex + 1; i < len(c.requests); i++ {
		if c.requests[i] == key {
			return i
		}
	}
	return -1
}

func main() {
	requests := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5,5,2,1,4,3,3,5,2,2,1,5,4,2,1,2,1,3}
	capacity := 3

	fmt.Println("Testing LRU Cache:")
	lru := NewLRUCache(capacity)
	lruMisses := 0
	for _, req := range requests {
		if !lru.Get(req) {
			fmt.Printf("LRU Cache miss for %d\n", req)
			lru.Put(req, req)
			lruMisses++
		}
	}
	fmt.Printf("Total LRU Cache misses: %d\n\n", lruMisses)

	fmt.Println("Testing Optimal Cache:")
	opt := NewOptimalCache(capacity, requests)
	optMisses := 0
	for i, req := range requests {
		if !opt.Get(req) {
			fmt.Printf("Optimal Cache miss for %d\n", req)
			opt.Put(req, req, i)
			optMisses++
		}
	}
	fmt.Printf("Total Optimal Cache misses: %d\n", optMisses)

	fmt.Printf("\n====\n")
	fmt.Printf("OPT - LRU = %d\n", lruMisses-optMisses)
}

/*
https://www.reddit.com/r/algorithms/comments/igkc1d/optimal_offline_caching_farthestinfuture_algorithm/
https://courses.cs.washington.edu/courses/cse421/18au/lecture/lecture-8.pdf
*/