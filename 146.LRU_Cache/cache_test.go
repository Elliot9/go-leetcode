package main

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		lRUCache := Constructor(2)
		lRUCache.Put(1, 1)
		lRUCache.Put(2, 2)
		if got := lRUCache.Get(1); got != 1 {
			t.Errorf("Get(1) = %v, want 1", got)
		}
		lRUCache.Put(3, 3)
		if got := lRUCache.Get(2); got != -1 {
			t.Errorf("Get(2) = %v, want -1", got)
		}
		lRUCache.Put(4, 4)
		if got := lRUCache.Get(1); got != -1 {
			t.Errorf("Get(1) = %v, want -1", got)
		}
		if got := lRUCache.Get(3); got != 3 {
			t.Errorf("Get(3) = %v, want 3", got)
		}
		if got := lRUCache.Get(4); got != 4 {
			t.Errorf("Get(4) = %v, want 4", got)
		}
	})

	t.Run("Capacity 1", func(t *testing.T) {
		lRUCache := Constructor(1)
		lRUCache.Put(2, 1)
		if got := lRUCache.Get(2); got != 1 {
			t.Errorf("Get(2) = %v, want 1", got)
		}
		lRUCache.Put(3, 2)
		if got := lRUCache.Get(2); got != -1 {
			t.Errorf("Get(2) = %v, want -1", got)
		}
		if got := lRUCache.Get(3); got != 2 {
			t.Errorf("Get(3) = %v, want 2", got)
		}
	})

	t.Run("Update Existing Key", func(t *testing.T) {
		lRUCache := Constructor(2)
		lRUCache.Put(1, 1)
		lRUCache.Put(2, 2)
		lRUCache.Put(1, 3)
		if got := lRUCache.Get(1); got != 3 {
			t.Errorf("Get(1) = %v, want 3", got)
		}
		if got := lRUCache.Get(2); got != 2 {
			t.Errorf("Get(2) = %v, want 2", got)
		}
	})

	t.Run("Get Non-existent Key", func(t *testing.T) {
		lRUCache := Constructor(2)
		if got := lRUCache.Get(1); got != -1 {
			t.Errorf("Get(1) = %v, want -1", got)
		}
	})
}
