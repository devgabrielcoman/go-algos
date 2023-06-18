package exercises_test

import (
	"gabriel/algos/exercises"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ArrayDuplication(t *testing.T) {
	t.Run("returns empty an empty array", func(t *testing.T) {
		array := []int{}
		result := exercises.ArrayDuplication(array)
		assert.Equal(t, []int{}, result)
	})

	t.Run("return empty for an array without duplication", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5}
		result := exercises.ArrayDuplication(array)
		assert.Equal(t, []int{}, result)
	})

	t.Run("return correct result for an array with duplication", func(t *testing.T) {
		array := []int{1, 2, 3, 3, 5}
		result := exercises.ArrayDuplication(array)
		assert.Equal(t, []int{3}, result)
	})
}
