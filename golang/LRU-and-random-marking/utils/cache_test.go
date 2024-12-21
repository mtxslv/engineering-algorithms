package utils

import (
	"reflect"
	"testing"
	"errors"
)

func TestBitwiseOperation(t *testing.T) {
	x := uint16(0) /* 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 */
	x = x + 1 << 5 /* 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 */
	if x != uint16(32){
		t.Fail()
	}
	i := 2
	x = x + 1 << i /* 0 0 0 0 0 0 0 0 0 0 0 1 0 1 0 0 */ 
	t.Logf("%d", x)
	if x != uint16(36) {
		t.Fail()
	}
}

func TestMakeMarkingCache(t *testing.T) {
	_, err := NewRandomMarkingCache(20) // 20 > 16
	if !errors.Is(err, ErrCacheSizeTooLarge) {
		t.Fail()
	}

	cache, err := NewRandomMarkingCache(15)
	if err != nil || cache.capacity != 15 || len(cache.cacheOrder) != 15{
		t.Fail()
	}
}

func TestCacheByte(t *testing.T) {
	cache, _ := NewRandomMarkingCache(10)
	cache.Mark(0) // Mark bit at 0
	cache.Mark(1) // Mark bit at 1
	if cache.markingByte != 3 {
		t.Fail()
	}
	
}

func TestCacheByteUnmark(t *testing.T) {
	cache, _ := NewRandomMarkingCache(16)
	i := 0
	for i < 5 {
		cache.Mark(i) /* 1 1 1 1 1 = 31 */
		i++
	}
	if cache.markingByte != 31 {
		t.Fail()
	}

	cache.Unmark(1)
	cache.Unmark(3)

	if cache.markingByte != 21 {
		t.Fail()
	}	
	i = 0
	for i < 16 {
		cache.Mark(i)
		i++
	}

	cache.Unmark(0)
	cache.Unmark(6)
	cache.Unmark(7)
	cache.Unmark(11)
	cache.Unmark(15)

	if cache.markingByte != 30526 {
		t.Fail()
	}

	i = 0
	for i < 16 {
		cache.Mark(i)
		i++
	}

	if cache.markingByte != 65535 { t.Fail() }

	i = 0
	for i < 16 {
		cache.Unmark(i)
		i++
	}
	if cache.markingByte != 0 { t.Fail() }

}

func TestMarkUnmarkLimit(t *testing.T) {
	cache, _ := NewRandomMarkingCache(2)
	err := cache.Mark(3)
	if !errors.Is(err, ErrAssignmentOutOfBounds) {
		t.Fail()
	}
	err = cache.Unmark(3)
	if !errors.Is(err, ErrAssignmentOutOfBounds) {
		t.Fail()
	}
}

func TestAllMarked(t *testing.T) {
	cache, _ := NewRandomMarkingCache(3)
	cache.Mark(0)
	cache.Mark(1)
	cache.Mark(2)
	if !cache.AllMarked(){
		t.Fail()
	}
}

func TestSelectUnmarkedButAllMarked(t *testing.T) {
	cache,_ := NewRandomMarkingCache(16)
	// MARK ALL BITS
	it := 0
	for it < 16 {
		cache.Mark(it)
		it++
	}
	// SELECT SHOULD RETURN -1
	if cache.SelectFromUnmarked() != -1 {
		t.Fail()
	}
}

func TestUnmarkedBits(t *testing.T) {
	cache,_ := NewRandomMarkingCache(4)
	cache.markingByte = 11
	unmarked := cache.UnmarkedBitsArray()
	expected := []int{2}
	t.Logf("%+v", unmarked)
	if !reflect.DeepEqual(unmarked, expected) {
		t.Fail()
	}
}