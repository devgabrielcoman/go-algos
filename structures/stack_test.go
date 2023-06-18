package structures_test

import (
	"gabriel/algos/structures"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Stack(t *testing.T) {
	t.Run("can create a new stack", func(t *testing.T) {
		stack := structures.Stack[int]()
		assert.NotNil(t, stack)
	})

	t.Run("can push to a stack", func(t *testing.T) {
		stack := structures.Stack[int]()
		assert.True(t, stack.IsEmpty())
		stack.Push(2)
		stack.Push(3)
		assert.False(t, stack.IsEmpty())
	})

	t.Run("can read a stack", func(t *testing.T) {
		stack := structures.Stack[int]()

		res1 := stack.Read()
		assert.Nil(t, res1)

		stack.Push(5)
		res2 := stack.Read()
		assert.Equal(t, 5, *res2)

		stack.Push(3)
		res3 := stack.Read()
		assert.Equal(t, 3, *res3)
	})

	t.Run("can pop a stack", func(t *testing.T) {
		stack := structures.Stack[int]()

		res1 := stack.Pop()
		assert.Nil(t, res1)

		stack.Push(5)
		res2 := stack.Pop()
		assert.Equal(t, 5, *res2)
		assert.True(t, stack.IsEmpty())

		stack.Push(5)
		stack.Push(3)
		assert.False(t, stack.IsEmpty())
		res3 := stack.Pop()
		res4 := stack.Pop()
		assert.True(t, stack.IsEmpty())
		assert.Equal(t, 3, *res3)
		assert.Equal(t, 5, *res4)
	})
}
