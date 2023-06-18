package algorithms

/**
 * We're progressively finding the lowest value in the array and then swapping that value with
 * the 0th, 1st, 2nd elements in the array.
 *
 * It runs in:
 * 	best case - O(N^2).
 * 	avg case - O(N^2)
 * 	worst case - O(N^2)
 *
 * input: 	7 1 5 2 0
 * 1st run: 0 1 5 2 7
 * 2nd run: 0 1 5 2 7
 * 3rd run: 0 1 2 5 7 --> sorted
 * 4th run: 0 1 2 5 7
 * 5th run: 0 1 2 5 7
 */
func SelectionSort(input []int) []int {
	result := make([]int, len(input))
	copy(result, input)

	for i := 0; i < len(result); i++ {
		var current_lowest_value_index = i

		for j := i + 1; j < len(result); j++ {
			if result[j] < result[current_lowest_value_index] {
				current_lowest_value_index = j
			}
		}

		if current_lowest_value_index != i {
			ith_value := result[i]
			lowest_value := result[current_lowest_value_index]
			result[current_lowest_value_index] = ith_value
			result[i] = lowest_value
		}
	}

	return result
}
