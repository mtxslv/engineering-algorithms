package utils

import (
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
	if err != nil || cache.capacity != 15 {
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