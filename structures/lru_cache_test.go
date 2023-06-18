package structures_test

import (
	"gabriel/algos/structures"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LRUCache(t *testing.T) {
	t.Run("can create a new lru cache", func(t *testing.T) {
		cache := structures.LRUCache(5)
		assert.NotNil(t, cache)
	})

	t.Run("will return nil if trying to get a value that does not exist", func(t *testing.T) {
		cache := structures.LRUCache(5)
		assert.Nil(t, cache.Get("key-1"))
	})

	t.Run("can insert and get upto capacity", func(t *testing.T) {
		cache := structures.LRUCache(3)
		cache.Put("key-1", 5)
		cache.Put("key-2", 3)
		cache.Put("key-3", 10)
		assert.Equal(t, 5, *cache.Get("key-1"))
		assert.Equal(t, 3, *cache.Get("key-2"))
		assert.Equal(t, 10, *cache.Get("key-3"))
	})

	t.Run("will start evicting data once it reaches capacity", func(t *testing.T) {
		cache := structures.LRUCache(3)
		cache.Put("key-1", 5)
		cache.Put("key-2", 3)
		cache.Put("key-3", 10)
		cache.Put("key-4", 14)
		assert.Nil(t, cache.Get("key-1"))
		assert.Equal(t, 3, *cache.Get("key-2"))
		assert.Equal(t, 10, *cache.Get("key-3"))
		assert.Equal(t, 14, *cache.Get("key-4"))
	})
}
