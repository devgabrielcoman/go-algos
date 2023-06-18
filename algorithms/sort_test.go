package algorithms_test

import (
	"gabriel/algos/algorithms"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Sort(t *testing.T) {
	empty_array := []int{}
	sorted_array := []int{1, 2, 3, 5, 7, 13, 21, 35, 65, 112, 256}
	reverse_array := []int{256, 112, 65, 35, 21, 13, 7, 5, 3, 2, 1}
	random_array := []int{1, 7, 3, 256, 112, 2, 21, 65, 35, 5, 13}
	negative_array := []int{-1, -7, -3, -256, -112, -2, -21, -65, -35, -5, -13}

	empty_result := []int{}
	sorted_result := []int{1, 2, 3, 5, 7, 13, 21, 35, 65, 112, 256}
	negative_result := []int{-256, -112, -65, -35, -21, -13, -7, -5, -3, -2, -1}

	t.Run("bubble sort returns sorted array in each case", func(t *testing.T) {
		result_1 := algorithms.BubbleSort(empty_array)
		assert.Equal(t, empty_result, result_1)

		result_2 := algorithms.BubbleSort(sorted_array)
		assert.Equal(t, sorted_result, result_2)

		result_3 := algorithms.BubbleSort(reverse_array)
		assert.Equal(t, sorted_result, result_3)

		result_4 := algorithms.BubbleSort(random_array)
		assert.Equal(t, sorted_result, result_4)

		result_5 := algorithms.BubbleSort(negative_array)
		assert.Equal(t, negative_result, result_5)
	})

	t.Run("insertion sort returns sorted array in each case", func(t *testing.T) {
		result_1 := algorithms.InsertionSort(empty_array)
		assert.Equal(t, empty_result, result_1)

		result_2 := algorithms.InsertionSort(sorted_array)
		assert.Equal(t, sorted_result, result_2)

		result_3 := algorithms.InsertionSort(reverse_array)
		assert.Equal(t, sorted_result, result_3)

		result_4 := algorithms.InsertionSort(random_array)
		assert.Equal(t, sorted_result, result_4)

		result_5 := algorithms.InsertionSort(negative_array)
		assert.Equal(t, negative_result, result_5)
	})

	t.Run("selection sort returns sorted array in each case", func(t *testing.T) {
		result_1 := algorithms.SelectionSort(empty_array)
		assert.Equal(t, empty_result, result_1)

		result_2 := algorithms.SelectionSort(sorted_array)
		assert.Equal(t, sorted_result, result_2)

		result_3 := algorithms.SelectionSort(reverse_array)
		assert.Equal(t, sorted_result, result_3)

		result_4 := algorithms.SelectionSort(random_array)
		assert.Equal(t, sorted_result, result_4)

		result_5 := algorithms.SelectionSort(negative_array)
		assert.Equal(t, negative_result, result_5)
	})

	t.Run("quick sort returns sorted array in each case", func(t *testing.T) {
		result_1 := algorithms.QuickSort(empty_array)
		assert.Equal(t, empty_result, result_1)

		result_2 := algorithms.QuickSort(sorted_array)
		assert.Equal(t, sorted_result, result_2)

		result_3 := algorithms.QuickSort(reverse_array)
		assert.Equal(t, sorted_result, result_3)

		result_4 := algorithms.QuickSort(random_array)
		assert.Equal(t, sorted_result, result_4)

		result_5 := algorithms.QuickSort(negative_array)
		assert.Equal(t, negative_result, result_5)
	})

	t.Run("binary sort returns sorted array in each case", func(t *testing.T) {
		result_1 := algorithms.BinarySort(empty_array)
		assert.Equal(t, empty_result, result_1)

		result_2 := algorithms.BinarySort(sorted_array)
		assert.Equal(t, sorted_result, result_2)

		result_3 := algorithms.BinarySort(reverse_array)
		assert.Equal(t, sorted_result, result_3)

		result_4 := algorithms.BinarySort(random_array)
		assert.Equal(t, sorted_result, result_4)

		result_5 := algorithms.BinarySort(negative_array)
		assert.Equal(t, negative_result, result_5)
	})
}
