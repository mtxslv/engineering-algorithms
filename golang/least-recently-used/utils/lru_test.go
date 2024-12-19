package utils

import (
	"testing"
)

func initializeEmptyCache() *LRUCacheV0{
	cacheSize := 3
	return NewLRUCacheV0(cacheSize)
}

func initializeEmptyCacheV1() *LRUCacheV1{
	cacheSize := 3
	return NewLRUCacheV1(cacheSize)
}

func initializeFullCache() *LRUCacheV0{
	cacheSize := 3
	lru := NewLRUCacheV0(cacheSize)
	
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


/////////////// LRU V1 ///////////////

func TestNewCacheV1(t *testing.T) {
	lru := initializeEmptyCacheV1()
	if lru.list.Len() != 0{
		t.Fail()
	}
}

func TestCacheV1Put(t *testing.T) {
	lru := initializeEmptyCacheV1()
	lru.Put("it",5.0)
	if lru.list.Len() != 1{
		t.Fail()
	}
	current := lru.list.Front()
	for current != nil {
		t.Logf("Value: %.3f", current.Value.(entryV1).value)
		current = current.Next()
	}
}

func TestCacheV1Update(t *testing.T){
	lru := initializeEmptyCacheV1()
	lru.Put("it",5.0) // it has original value 5.0
	lru.Put("it",13.47) // now it has value 13.47
	if lru.list.Front().Value.(entryV1).value != 13.47 {
		t.Fail()
	}
}

func TestCacheGet(t *testing.T){
	lru := initializeEmptyCacheV1()
	lru.Put("it",5.0) // it has original value 5.0
	val, _ := lru.Get("it") // Get value
	if val != 5.0 {
		t.Fail()
	}	
	_, found := lru.Get("doesnotexist") 
	if found == true {
		t.Fail()
	}
}

////////// OPT CACHE //////////////
var requests = []string{
	"rick", "morty", "rick", "morty", "summer", 
	"beth", "rick", "morty", "rick", "rick", 
	"beth", "jerry",
}
var data = map[string]float32{
	"rick":70.0,
	"beth":34.0,
	"jerry":35.0,
	"summer":17.0,
	"morty":14.0,
}

func initializeEmptyCacheOPT() *LRUCacheOPT{
	cacheSize := 3
	return NewOPTCache(cacheSize, requests)
}

func TestCacheOptWhoToEvict(t *testing.T){
	
	lruOpt := initializeEmptyCacheOPT()
	it := 0
	for it < 5 {
		lruOpt.Put(
			requests[it],
			data[requests[it]],
		)
		lruOpt.lastRequestNum++
		it++
	}
	
	// current := lruOpt.list.Front()
	// for current != nil {
	// 	t.Logf(
	// 		"Age of %s : %.2f years", 
	// 		current.Value.(entryV1).key,
	// 		current.Value.(entryV1).value,
	// 	)
	// 	current = current.Next()
	// }
	// t.Logf("LAST QUERY MADE? %d", lruOpt.lastRequestNum)
	toEvict := lruOpt.CheckLateUse()
	if toEvict.Value.(entryV1).key != "summer" {
		t.Fail()
	}
	// t.Logf(
	// 	"Evict %s ",
	// 	toEvict.Value.(entryV1).key,
	// )
}