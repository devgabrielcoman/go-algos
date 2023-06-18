package exercises_test

import (
	"gabriel/algos/exercises"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ArrayIntersection(t *testing.T) {
	t.Run("returns empty for the interesction of two empty arrays", func(t *testing.T) {
		array1 := []int{}
		array2 := []int{}
		result := exercises.ArrayIntersection(array1, array2)
		assert.Equal(t, []int{}, result)
	})

	t.Run("returns empty for the interesction of two disjoing arrays", func(t *testing.T) {
		array1 := []int{1, 2, 3, 4, 5}
		array2 := []int{6, 7, 8, 9, 10}
		result := exercises.ArrayIntersection(array1, array2)
		assert.Equal(t, []int{}, result)
	})

	t.Run("returns correct result for the interesction of two arrays", func(t *testing.T) {
		array1 := []int{1, 2, 3, 4, 5}
		array2 := []int{3, 4, 5, 6, 7}
		result := exercises.ArrayIntersection(array1, array2)
		assert.Equal(t, []int{3, 4, 5}, result)
	})

	t.Run("returns same array for the interesction of two identical arrays", func(t *testing.T) {
		array1 := []int{1, 2, 3, 4, 5}
		array2 := []int{1, 2, 3, 4, 5}
		result := exercises.ArrayIntersection(array1, array2)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
	})
}
