package structures

type stack[T any] struct {
	data []T
}

func Stack[T any]() *stack[T] {
	return &stack[T]{
		data: []T{},
	}
}

func (s *stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

func (s *stack[T]) Pop() *T {
	element_to_pop := s.Read()

	if element_to_pop == nil {
		return nil
	}

	s.data = s.data[0 : len(s.data)-1]

	return element_to_pop
}

func (s *stack[T]) Read() *T {
	current_len := len(s.data)

	if current_len == 0 {
		return nil
	}

	return &s.data[current_len-1]
}
