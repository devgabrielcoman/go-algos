package structures

type singleNode[T any] struct {
	Value T
	next  *singleNode[T]
}

func newSingleNode[T any](Value T, next *singleNode[T]) *singleNode[T] {
	return &singleNode[T]{Value, next}
}

type linkedList[T any] struct {
	start *singleNode[T]
}

func LinkedList[T any]() *linkedList[T] {
	return &linkedList[T]{
		start: nil,
	}
}

func (l *linkedList[T]) IsEmpty() bool {
	return l.start == nil
}

func (l *linkedList[T]) AddAtStart(value T) {
	next := newSingleNode(value, l.start)
	l.start = next
}

func (l *linkedList[T]) RemoveAtStart() *singleNode[T] {
	removed := l.start
	l.start = l.safeGetNextNode(l.start)
	return removed
}

func (l *linkedList[T]) AddAtEnd(value T) {
	last, _ := l.getLastNodes()
	next := newSingleNode(value, nil)
	l.safeAssignAfter(last, next)
}

func (l *linkedList[T]) RemoveAtEnd() *singleNode[T] {
	last, previous := l.getLastNodes()
	l.safeRemoveNext(previous)
	return last
}

func (l *linkedList[T]) ToArray() []T {
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

func (l *linkedList[T]) Length() int {
	return len(l.ToArray())
}

func (l *linkedList[T]) GetEnd() *singleNode[T] {
	last, _ := l.getLastNodes()
	return last
}

func (l *linkedList[T]) GetStart() *singleNode[T] {
	return l.start
}

func (l *linkedList[T]) AddAtIndex(index int, value T) {
	if index < 0 {
		return
	}

	if index == 0 {
		l.AddAtStart(value)
		return
	}

	var current = l.start
	for i := 0; i < index-1 && l.hasNextNode(current); i++ {
		current = l.safeGetNextNode(current)
	}
	next := newSingleNode(value, l.safeGetNextNode(current))
	l.safeAssignAfter(current, next)
}

func (l *linkedList[T]) RemoveAtIndex(index int) *singleNode[T] {
	if index < 0 {
		return nil
	}

	if index == 0 {
		return l.RemoveAtStart()
	}

	var previous = l.start
	for i := 0; i < index-1 && l.hasNextNode(previous); i++ {
		previous = l.safeGetNextNode(previous)
	}

	tmp_node := previous.next
	previous.next = l.safeGetNextNode(previous.next)
	return tmp_node
}

func (l *linkedList[T]) getLastNodes() (*singleNode[T], *singleNode[T]) {
	var current = l.start
	var prev *singleNode[T] = nil
	for l.hasNextNode(current) {
		prev = current
		current = l.safeGetNextNode(current)
	}
	return current, prev
}

func (l *linkedList[T]) safeAssignAfter(current *singleNode[T], next *singleNode[T]) {
	if current == nil {
		l.start = next
	} else {
		current.next = next
	}
}

func (l *linkedList[T]) safeRemoveNext(current *singleNode[T]) {
	if current == nil {
		l.start = nil
	} else {
		current.next = nil
	}
}

func (l *linkedList[T]) hasNextNode(current *singleNode[T]) bool {
	return current != nil && current.next != nil
}

func (l *linkedList[T]) safeGetNextNode(current *singleNode[T]) *singleNode[T] {
	if current == nil {
		return nil
	}
	return current.next
}
