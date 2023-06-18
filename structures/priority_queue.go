package structures

import "golang.org/x/exp/constraints"

type proprityQueue[T constraints.Ordered] struct {
	heap *heap[T]
}

func PriorityQueue[T constraints.Ordered]() *proprityQueue[T] {
	return &proprityQueue[T]{
		heap: Heap[T](),
	}
}

func (q *proprityQueue[T]) IsEmpty() bool {
	return q.heap.IsEmpty()
}

func (q *proprityQueue[T]) Enqueue(value T) {
	q.heap.Insert(value)
}

func (q *proprityQueue[T]) Dequeue() *T {
	return q.heap.Delete()
}

func (q *proprityQueue[T]) Read() *T {
	return q.heap.ReadFirst()
}
