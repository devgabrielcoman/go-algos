package algorithms_test

import (
	"gabriel/algos/algorithms"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BinarySearch(t *testing.T) {
	t.Run("returns false if searching through an empty array", func(t *testing.T) {
		array := []int{}
		result := algorithms.BinarySearch(array, 18)
		assert.False(t, result)
	})

	t.Run("returns false if searching through a one element array that does not contain element", func(t *testing.T) {
		array := []int{3}
		result := algorithms.BinarySearch(array, 18)
		assert.False(t, result)
	})

	t.Run("returns true if searching through a one element array containing the exact element", func(t *testing.T) {
		array := []int{18}
		result := algorithms.BinarySearch(array, 18)
		assert.True(t, result)
	})

	t.Run("returns true if it can find the desired value in an even numbered ordered array", func(t *testing.T) {
		array := []int{1, 2, 3, 5, 7, 13, 21, 25, 27, 48, 53, 58, 61, 62, 63, 71, 75, 78, 82, 85, 87, 90, 91, 92, 95, 97, 101, 102}
		result := algorithms.BinarySearch(array, 95)
		assert.True(t, result)
	})

	t.Run("returns true if it can find the desired value in an odd numbered ordered array", func(t *testing.T) {
		array := []int{2, 3, 5, 7, 13, 21, 25, 27, 48, 53, 58, 61, 62, 63, 71, 75, 78, 82, 85, 87, 90, 91, 92, 95, 97, 101, 102}
		result := algorithms.BinarySearch(array, 95)
		assert.True(t, result)
	})

	t.Run("returns false if it cannot find the desired value in an even numbered ordered array", func(t *testing.T) {
		array := []int{1, 2, 3, 5, 7, 13, 21, 25, 27, 48, 53, 58, 61, 62, 63, 71, 75, 78, 82, 85, 87, 90, 91, 92, 95, 97, 101, 102}
		result := algorithms.BinarySearch(array, 18)
		assert.False(t, result)
	})

	t.Run("returns false if it cannot find the desired value in an odd numbered ordered array", func(t *testing.T) {
		array := []int{2, 3, 5, 7, 13, 21, 25, 27, 48, 53, 58, 61, 62, 63, 71, 75, 78, 82, 85, 87, 90, 91, 92, 95, 97, 101, 102}
		result := algorithms.BinarySearch(array, 18)
		assert.False(t, result)
	})
}
