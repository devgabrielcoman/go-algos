package structures

import (
	"golang.org/x/exp/constraints"
)

/**
 * A binary tree that maintains two conditions:
 * - The value of each node must be greater than that of each of its descendant nodes
 * - The tree must be complete (from left-to-right, on each level, all nodes are there; bottom row can have empty positions, as long as there aren't any other right-er nodes)
 * Insertions and deletions are both O(logN), compared with ordered arrays, which are O(N) and O(1) respectively.
 */
type heap[T constraints.Ordered] struct {
	data []T
}

func Heap[T constraints.Ordered]() *heap[T] {
	return &heap[T]{
		data: []T{},
	}
}

func (h *heap[T]) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *heap[T]) ToArray() []T {
	return h.data
}

func (h *heap[T]) ReadFirst() *T {
	if h.IsEmpty() {
		return nil
	}

	return &h.data[0]
}

func (h *heap[T]) Insert(value T) {
	// add the node as the "last node"
	h.data = append(h.data, value)
	// then we perform the "trickle up" operation
	last_node_index := len(h.data) - 1
	h.trickle_up(last_node_index)
}

func (h *heap[T]) Delete() *T {
	if h.IsEmpty() {
		return nil
	}

	last_node_index := len(h.data) - 1
	root_node_index := 0
	root_node_value := h.data[root_node_index]

	// swap last node with root node
	h.swap(root_node_index, last_node_index)
	h.data = h.data[0 : len(h.data)-1]

	if len(h.data) > 0 {
		h.trickle_down(root_node_index)
	}

	return &root_node_value
}

func (h *heap[T]) trickle_up(last_node_index int) {
	parent_index := h.parent_index(last_node_index)
	parent_index_value := h.data[parent_index]
	last_node_index_value := h.data[last_node_index]

	if last_node_index_value > parent_index_value {
		h.swap(parent_index, last_node_index)
		h.trickle_up(parent_index)
	}
}

func (h *heap[T]) trickle_down(trickle_node_index int) {
	trickle_node_value := h.data[trickle_node_index]

	left_child_index := h.left_child_index(trickle_node_index)
	right_child_index := h.right_child_index(trickle_node_index)

	if left_child_index >= len(h.data) {
		return // end of array
	}

	left_child_value := h.data[left_child_index]

	var right_child_value = left_child_value
	if right_child_index < len(h.data) {
		right_child_value = h.data[right_child_index]
	}

	if left_child_value >= right_child_value {
		if trickle_node_value < left_child_value {
			h.swap(trickle_node_index, left_child_index)
			h.trickle_down(left_child_index)
		}
	} else {
		if trickle_node_value < right_child_value {
			h.swap(trickle_node_index, right_child_index)
			h.trickle_down(right_child_index)
		}
	}
}

func (h *heap[T]) left_child_index(index int) int {
	return (index * 2) + 1
}

func (h *heap[T]) right_child_index(index int) int {
	return (index * 2) + 2
}

func (h *heap[T]) parent_index(index int) int {
	return (index - 1) / 2
}

func (h *heap[T]) swap(i int, j int) {
	value_at_j := h.data[j]
	h.data[j] = h.data[i]
	h.data[i] = value_at_j
}
