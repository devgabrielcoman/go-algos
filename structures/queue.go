package structures

type queue[T any] struct {
	list *doublyLinkedList[T]
}

func Queue[T any]() *queue[T] {
	return &queue[T]{
		list: DoublyLinkedList[T](),
	}
}

func (q *queue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *queue[T]) Enqueue(value T) {
	q.list.AddAtEnd(value)
}

func (q *queue[T]) Dequeue() *T {
	node := q.list.RemoveAtStart()
	return q.readSafeVal(node)
}

func (q *queue[T]) Read() *T {
	node := q.list.GetStart()
	return q.readSafeVal(node)
}

func (q *queue[T]) readSafeVal(node *doubleNode[T]) *T {
	if node != nil {
		return &node.Value
	} else {
		return nil
	}
}
