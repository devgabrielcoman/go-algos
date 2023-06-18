package structures_test

import (
	"gabriel/algos/structures"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DoublyLinkedList(t *testing.T) {
	t.Run("can create a new doubly linked list", func(t *testing.T) {
		linkedList := structures.DoublyLinkedList[int]()
		assert.NotNil(t, linkedList)
	})

	t.Run("isEmpty returns true for a newly created linked list", func(t *testing.T) {
		linkedList := structures.DoublyLinkedList[int]()
		assert.True(t, linkedList.IsEmpty())
	})

	t.Run("can add at start", func(t *testing.T) {
		linkedList := structures.DoublyLinkedList[int]()
		linkedList.AddAtStart(5)
		assert.False(t, linkedList.IsEmpty())
		linkedList.AddAtStart(3)
		assert.False(t, linkedList.IsEmpty())
		array := linkedList.ToArray()
		assert.Equal(t, []int{3, 5}, array)
	})

	t.Run("can remove at start", func(t *testing.T) {
		linkedList := structures.DoublyLinkedList[int]()
		linkedList.AddAtStart(5)
		linkedList.AddAtStart(3)
		linkedList.AddAtStart(8)
		assert.Equal(t, []int{8, 3, 5}, linkedList.ToArray())
		node1 := linkedList.RemoveAtStart()
		assert.Equal(t, 8, node1.Value)
		assert.Equal(t, []int{3, 5}, linkedList.ToArray())
		node2 := linkedList.RemoveAtStart()
		assert.Equal(t, 3, node2.Value)
		assert.Equal(t, []int{5}, linkedList.ToArray())
		node3 := linkedList.RemoveAtStart()
		assert.Equal(t, 5, node3.Value)
		assert.Equal(t, []int{}, linkedList.ToArray())
		node4 := linkedList.RemoveAtStart()
		assert.Nil(t, node4)
		assert.Equal(t, []int{}, linkedList.ToArray())
	})

	t.Run("can add at end", func(t *testing.T) {
		linkedList := structures.DoublyLinkedList[int]()
		linkedList.AddAtEnd(5)
		assert.False(t, linkedList.IsEmpty())
		linkedList.AddAtEnd(3)
		assert.False(t, linkedList.IsEmpty())
		array := linkedList.ToArray()
		assert.Equal(t, []int{5, 3}, array)
	})

	t.Run("can remove at end", func(t *testing.T) {
		linkedList := structures.DoublyLinkedList[int]()
		linkedList.AddAtEnd(5)
		linkedList.AddAtEnd(8)
		linkedList.AddAtEnd(12)
		assert.Equal(t, []int{5, 8, 12}, linkedList.ToArray())
		node1 := linkedList.RemoveAtEnd()
		assert.Equal(t, 12, node1.Value)
		assert.Equal(t, []int{5, 8}, linkedList.ToArray())
		node2 := linkedList.RemoveAtEnd()
		assert.Equal(t, 8, node2.Value)
		assert.Equal(t, []int{5}, linkedList.ToArray())
		node3 := linkedList.RemoveAtEnd()
		assert.Equal(t, 5, node3.Value)
		assert.Equal(t, []int{}, linkedList.ToArray())
		node4 := linkedList.RemoveAtEnd()
		assert.Nil(t, node4)
		assert.Equal(t, []int{}, linkedList.ToArray())
	})

	t.Run("can add at start and end", func(t *testing.T) {
		linkedList := structures.DoublyLinkedList[int]()
		linkedList.AddAtEnd(5)
		linkedList.AddAtEnd(3)
		linkedList.AddAtStart(7)
		array := linkedList.ToArray()
		assert.Equal(t, []int{7, 5, 3}, array)
	})

	t.Run("can remove first element", func(t *testing.T) {
		linkedList := structures.DoublyLinkedList[int]()
		res1 := linkedList.RemoveAtStart()
		assert.Nil(t, res1)
		linkedList.AddAtStart(5)
		linkedList.AddAtStart(3)
		linkedList.AddAtStart(7)
		res2 := linkedList.RemoveAtStart()
		assert.Equal(t, 7, res2.Value)
		array1 := linkedList.ToArray()
		assert.Equal(t, []int{3, 5}, array1)
		res3 := linkedList.RemoveAtStart()
		assert.Equal(t, 3, res3.Value)
		array2 := linkedList.ToArray()
		assert.Equal(t, []int{5}, array2)
		res4 := linkedList.RemoveAtStart()
		assert.Equal(t, 5, res4.Value)
		array3 := linkedList.ToArray()
		assert.Equal(t, []int{}, array3)
		assert.True(t, linkedList.IsEmpty())
	})

	t.Run("can count the number of elements in the list", func(t *testing.T) {
		linkedList := structures.DoublyLinkedList[int]()
		assert.Equal(t, 0, linkedList.Length())
		linkedList.AddAtStart(7)
		linkedList.AddAtEnd(7)
		linkedList.AddAtStart(3)
		assert.Equal(t, 3, linkedList.Length())
	})

	t.Run("can read the start", func(t *testing.T) {
		linkedList := structures.DoublyLinkedList[int]()
		res1 := linkedList.GetStart()
		assert.Nil(t, res1)

		linkedList.AddAtStart(5)
		res2 := linkedList.GetStart()
		assert.Equal(t, 5, res2.Value)

		linkedList.AddAtStart(8)
		res3 := linkedList.GetStart()
		assert.Equal(t, 8, res3.Value)

		linkedList.AddAtEnd(12)
		res4 := linkedList.GetStart()
		assert.Equal(t, 8, res4.Value)
	})

	t.Run("can read the end", func(t *testing.T) {
		linkedList := structures.DoublyLinkedList[int]()
		res1 := linkedList.GetEnd()
		assert.Nil(t, res1)

		linkedList.AddAtStart(5)
		res2 := linkedList.GetEnd()
		assert.Equal(t, 5, res2.Value)

		linkedList.AddAtStart(8)
		res3 := linkedList.GetEnd()
		assert.Equal(t, 5, res3.Value)

		linkedList.AddAtEnd(12)
		res4 := linkedList.GetEnd()
		assert.Equal(t, 12, res4.Value)
	})
}
