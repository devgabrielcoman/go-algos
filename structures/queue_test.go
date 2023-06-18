package structures_test

import (
	"gabriel/algos/structures"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Queue(t *testing.T) {
	t.Run("can create a new queue", func(t *testing.T) {
		queue := structures.Queue[int]()
		assert.NotNil(t, queue)
	})

	t.Run("a queue with no elements will always be empty", func(t *testing.T) {
		queue := structures.Queue[int]()
		assert.True(t, queue.IsEmpty())
	})

	t.Run("can enqueue elements", func(t *testing.T) {
		queue := structures.Queue[int]()
		queue.Enqueue(5)
		assert.Equal(t, 5, *queue.Read())
		queue.Enqueue(7)
		assert.Equal(t, 5, *queue.Read())
		queue.Enqueue(13)
		assert.Equal(t, 5, *queue.Read())
		assert.False(t, queue.IsEmpty())
	})

	t.Run("will return nill when reading an empty queue", func(t *testing.T) {
		queue := structures.Queue[int]()
		assert.Nil(t, queue.Read())
	})

	t.Run("can dequeue elements", func(t *testing.T) {
		queue := structures.Queue[int]()

		res1 := queue.Dequeue()
		assert.Nil(t, res1)

		queue.Enqueue(5)
		res2 := queue.Dequeue()
		assert.Equal(t, 5, *res2)
		assert.True(t, queue.IsEmpty())

		queue.Enqueue(5)
		queue.Enqueue(7)
		queue.Enqueue(13)
		assert.False(t, queue.IsEmpty())

		res3 := queue.Dequeue()
		res4 := queue.Dequeue()
		res5 := queue.Dequeue()
		assert.Equal(t, 5, *res3)
		assert.Equal(t, 7, *res4)
		assert.Equal(t, 13, *res5)
		assert.True(t, queue.IsEmpty())
	})
}
