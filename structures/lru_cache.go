package structures

/**
 * A LRU cache of a certain capacity will hold N elements.
 * Once it reaches capacity, it will evict the oldest element to make room.
 * We assume capacity is positive.
 * We use a DoublyLinked list to achieve O(1) insertions and deletions
 */
type lruCache struct {
	capacity  int
	list      *doublyLinkedList[string]
	container map[string]int
}

func LRUCache(capacity int) *lruCache {
	return &lruCache{
		capacity:  capacity,
		list:      DoublyLinkedList[string](),
		container: make(map[string]int),
	}
}

func (l *lruCache) Put(key string, value int) {
	// we still have room
	if len(l.container) < l.capacity {
		l.list.AddAtStart(key)
	} else {
		// evict the last element
		last_key := l.list.RemoveAtEnd().Value
		delete(l.container, last_key)
	}

	l.container[key] = value
}

func (l *lruCache) Get(key string) *int {
	value, exists := l.container[key]
	if !exists {
		return nil
	}

	return &value
}
