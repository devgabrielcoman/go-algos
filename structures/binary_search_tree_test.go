package structures_test

import (
	"gabriel/algos/structures"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BinarySearchTree(t *testing.T) {
	t.Run("can create a new bst", func(t *testing.T) {
		bst := structures.BinarySearchTree[int]()
		assert.NotNil(t, bst)
	})

	t.Run("can insert a new element", func(t *testing.T) {
		bst := structures.BinarySearchTree[int]()
		assert.True(t, bst.IsEmpty())
		assert.Equal(t, 0, bst.Length())
		bst.Insert(5)
		assert.False(t, bst.IsEmpty())
		assert.Equal(t, 1, bst.Length())
		assert.Equal(t, []int{5}, bst.ToArray())
		bst.Insert(7)
		bst.Insert(3)
		bst.Insert(12)
		assert.Equal(t, 4, bst.Length())
		assert.Equal(t, []int{3, 5, 7, 12}, bst.ToArray())
	})
}
