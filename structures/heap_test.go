package structures_test

import (
	"gabriel/algos/structures"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Heap(t *testing.T) {
	t.Run("can create a new heap", func(t *testing.T) {
		heap := structures.Heap[int]()
		assert.NotNil(t, heap)
	})

	t.Run("can correctly insert a new element", func(t *testing.T) {
		heap := structures.Heap[int]()
		assert.True(t, heap.IsEmpty())
		assert.Nil(t, heap.ReadFirst())
		heap.Insert(5)
		assert.Equal(t, 5, *heap.ReadFirst())
		heap.Insert(7)
		assert.Equal(t, 7, *heap.ReadFirst())
		heap.Insert(3)
		assert.Equal(t, 7, *heap.ReadFirst())
		heap.Insert(21)
		assert.Equal(t, 21, *heap.ReadFirst())
		heap.Insert(-1)
		assert.Equal(t, 21, *heap.ReadFirst())
	})

	t.Run("can correctly delete the first element", func(t *testing.T) {
		heap := structures.Heap[int]()
		assert.True(t, heap.IsEmpty())
		assert.Nil(t, heap.ReadFirst())
		assert.Nil(t, heap.Delete())

		heap.Insert(5)
		heap.Insert(7)
		heap.Insert(3)
		heap.Insert(21)
		heap.Insert(-1)

		assert.Equal(t, 21, *heap.ReadFirst())

		assert.Equal(t, 21, *heap.Delete())
		assert.Equal(t, 7, *heap.Delete())
		assert.Equal(t, 5, *heap.Delete())
		assert.Equal(t, 3, *heap.Delete())
		assert.Equal(t, -1, *heap.Delete())
		assert.Nil(t, heap.Delete())
	})

	t.Run("can correctly insert and delete", func(t *testing.T) {
		heap := structures.Heap[int]()

		heap.Insert(5)
		assert.Equal(t, 5, *heap.ReadFirst())

		heap.Insert(-1)
		heap.Insert(21)
		assert.Equal(t, 21, *heap.ReadFirst())

		assert.Equal(t, 21, *heap.Delete())
		assert.Equal(t, 5, *heap.ReadFirst())
		assert.Equal(t, 5, *heap.Delete())
		assert.Equal(t, -1, *heap.ReadFirst())

		heap.Insert(15)
		heap.Insert(-3)
		heap.Insert(7)
		heap.Insert(14)

		assert.Equal(t, 15, *heap.ReadFirst())
		assert.Equal(t, 15, *heap.Delete())
		assert.Equal(t, 14, *heap.Delete())
		assert.Equal(t, 7, *heap.ReadFirst())
	})
}
