package structures

type doubleNode[T any] struct {
	Value T
	next  *doubleNode[T]
	prev  *doubleNode[T]
}

func newDoubleNode[T any](Value T) *doubleNode[T] {
	return &doubleNode[T]{Value: Value, next: nil, prev: nil}
}

type doublyLinkedList[T any] struct {
	start *doubleNode[T]
	end   *doubleNode[T]
}

func DoublyLinkedList[T any]() *doublyLinkedList[T] {
	return &doublyLinkedList[T]{
		start: nil,
		end:   nil,
	}
}

func (l *doublyLinkedList[T]) IsEmpty() bool {
	return l.start == nil
}

func (l *doublyLinkedList[T]) AddAtStart(value T) {
	node := newDoubleNode(value)

	if l.IsEmpty() {
		l.start = node
		l.end = node
	} else {
		l.start.prev = node
		node.next = l.start
		l.start = node
	}
}

func (l *doublyLinkedList[T]) RemoveAtStart() *doubleNode[T] {
	if l.IsEmpty() {
		return nil
	}

	removed := l.start
	next := l.start.next
	if next != nil {
		next.prev = nil
	} else {
		l.end = nil
	}
	l.start = next

	return removed
}

func (l *doublyLinkedList[T]) AddAtEnd(value T) {
	node := newDoubleNode(value)

	if l.IsEmpty() {
		l.start = node
		l.end = node
	} else {
		l.end.next = node
		node.prev = l.end
		l.end = node
	}
}

func (l *doublyLinkedList[T]) RemoveAtEnd() *doubleNode[T] {
	if l.IsEmpty() {
		return nil
	}

	removed := l.end
	prev := l.end.prev
	if prev != nil {
		prev.next = nil
	} else {
		l.start = nil
	}
	l.end = prev

	return removed
}

func (l *doublyLinkedList[T]) ToArray() []T {
	var result = []T{}

	var current = l.start
	if current != nil {
		result = append(result, current.Value)
	}

	for l.hasNextNode(current) {
		current = l.safeGetNextNode(current)
		result = append(result, current.Value)
	}

	return result
}

func (l *doublyLinkedList[T]) Length() int {
	return len(l.ToArray())
}

func (l *doublyLinkedList[T]) GetEnd() *doubleNode[T] {
	return l.end
}

func (l *doublyLinkedList[T]) GetStart() *doubleNode[T] {
	return l.start
}

func (l *doublyLinkedList[T]) hasNextNode(current *doubleNode[T]) bool {
	return current != nil && current.next != nil
}

func (l *doublyLinkedList[T]) safeGetNextNode(current *doubleNode[T]) *doubleNode[T] {
	if current == nil {
		return nil
	}
	return current.next
}
