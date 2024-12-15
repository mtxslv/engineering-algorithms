package utils

import (
	"testing"
)

func initializeEmptyCache() *LRUCache{
	cacheSize := 3
	return NewLRUCache(cacheSize)
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