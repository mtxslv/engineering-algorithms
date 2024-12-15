package utils

import (
	"testing"
)

func initializeEmptyCache() *LRUCache{
	cacheSize := 3
	return NewLRUCache(cacheSize)
}

func initializeFullCache() *LRUCache{
	cacheSize := 3
	lru := NewLRUCache(cacheSize)
	
	lru.Put(1,10)
	lru.Put(2,20)
	lru.Put(3,30)

	return lru
}

func TestPut(t *testing.T){
	lruCache := initializeEmptyCache()
	lruCache.Put(5,50)
	lruCache.Put(1,10)
		
	for k := range lruCache.cache{
		el := lruCache.cache[k]
		t.Logf("AN ITEM: %+v\n", el.Value)
	}

	lruCache.Put(1,100)
	t.Logf("REPLACE ELEMENT {1,10} by {1,100}")
	for k := range lruCache.cache{
		el := lruCache.cache[k]
		t.Logf("AN ITEM: %+v\n", el.Value)
	}
}

func TestGet(t *testing.T){
	lru := initializeFullCache()
	current := lru.list.Front()
	// items are:
	// {3:30}->{2:20}->{1:10}
	for current != nil {
		t.Logf("LIST ITEM: %+v", current.Value)
		current = current.Next()
	}

	t.Logf("Search for 1\n")
	lru.Get(1)
	
	// Start from the beginning
	current = lru.list.Front()
	for current != nil {
		t.Logf("LIST ITEM: %+v", current.Value)
		current = current.Next()
	}
	
}

func TestPutNewTerm(t *testing.T){
	lru := initializeFullCache()

	// Start from the beginning
	current := lru.list.Front()
	for current != nil {
		t.Logf("LIST ITEM: %+v", current.Value)
		current = current.Next()
	}
	
	lru.Put(4,40)
	current = lru.list.Front()
	for current != nil {
		t.Logf("LIST ITEM: %+v", current.Value)
		current = current.Next()
	}
}