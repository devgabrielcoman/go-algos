package structures

type hashdatum[T comparable] struct {
	key   string
	value T
}

type hasharray[T comparable] struct {
	array []hashdatum[T]
}

func newHashArray[T comparable]() hasharray[T] {
	return hasharray[T]{
		array: []hashdatum[T]{},
	}
}

type hashmap[T comparable] struct {
	data [9]hasharray[T]
}

func HashMap[T comparable]() *hashmap[T] {
	data := [9]hasharray[T]{}
	for i := 0; i < 9; i++ {
		data[i] = newHashArray[T]()
	}

	return &hashmap[T]{data}
}

func (s *hashmap[T]) Set(key string, value T) {
	index := s.Hash(key)
	item := s.data[index]
	current_array := item.array

	for i, datum := range current_array {
		// if we've found an element with the same key
		if datum.key == key {
			// if the values differ, overwrite it
			if datum.value != value {
				current_array[i].value = value
			}
			return
		}
	}

	// we haven't found an existing element to overwrite or leave as is
	datum := hashdatum[T]{key: key, value: value}
	new_array := append(current_array, datum)
	item.array = new_array
	s.data[index] = item
}

func (s *hashmap[T]) Get(key string) *T {
	index := s.Hash(key)
	item := s.data[index]

	for _, value := range item.array {
		if value.key == key {
			return &value.value
		}
	}

	return nil
}

func (s *hashmap[T]) Hash(key string) int {
	hash := 0

	for _, char := range key {
		hash += int(char)
	}

	// Apply modulo operation to restrict the hash value between 1 and 9
	hash = (hash % 9) + 1

	return hash
}
