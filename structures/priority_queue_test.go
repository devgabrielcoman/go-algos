package structures_test

import (
	"gabriel/algos/structures"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PriorityQueue(t *testing.T) {
	t.Run("can create a priority queue", func(t *testing.T) {
		queue := structures.PriorityQueue[int]()
		assert.NotNil(t, queue)
	})

	t.Run("can correctly insert a new element", func(t *testing.T) {
		queue := structures.PriorityQueue[int]()
		assert.True(t, queue.IsEmpty())
		assert.Nil(t, queue.Read())
		queue.Enqueue(5)
		assert.Equal(t, 5, *queue.Read())
		queue.Enqueue(7)
		assert.Equal(t, 7, *queue.Read())
		queue.Enqueue(3)
		assert.Equal(t, 7, *queue.Read())
		queue.Enqueue(21)
		assert.Equal(t, 21, *queue.Read())
		queue.Enqueue(-1)
		assert.Equal(t, 21, *queue.Read())
	})

	t.Run("can correctly delete the first element", func(t *testing.T) {
		queue := structures.PriorityQueue[int]()
		assert.True(t, queue.IsEmpty())
		assert.Nil(t, queue.Read())
		assert.Nil(t, queue.Dequeue())

		queue.Enqueue(5)
		queue.Enqueue(7)
		queue.Enqueue(3)
		queue.Enqueue(21)
		queue.Enqueue(-1)

		assert.Equal(t, 21, *queue.Read())

		assert.Equal(t, 21, *queue.Dequeue())
		assert.Equal(t, 7, *queue.Dequeue())
		assert.Equal(t, 5, *queue.Dequeue())
		assert.Equal(t, 3, *queue.Dequeue())
		assert.Equal(t, -1, *queue.Dequeue())
		assert.Nil(t, queue.Dequeue())
	})

	t.Run("can correctly insert and delete", func(t *testing.T) {
		queue := structures.PriorityQueue[int]()

		queue.Enqueue(5)
		assert.Equal(t, 5, *queue.Read())

		queue.Enqueue(-1)
		queue.Enqueue(21)
		assert.Equal(t, 21, *queue.Read())

		assert.Equal(t, 21, *queue.Dequeue())
		assert.Equal(t, 5, *queue.Read())
		assert.Equal(t, 5, *queue.Dequeue())
		assert.Equal(t, -1, *queue.Read())

		queue.Enqueue(15)
		queue.Enqueue(-3)
		queue.Enqueue(7)
		queue.Enqueue(14)

		assert.Equal(t, 15, *queue.Read())
		assert.Equal(t, 15, *queue.Dequeue())
		assert.Equal(t, 14, *queue.Dequeue())
		assert.Equal(t, 7, *queue.Read())
	})
}
