package algorithms

/**
 * We temporary remove the value at index 1, 2, ... N and we do a comparison with all the remaining left-wise positons.
 * If the left-wise position is greater than the temporary value, we shift the value to the right.
 * We do that until the end.
 *
 * It runs in:
 * 	best case - O(N).
 * 	avg case - O(N^2)
 * 	worst case - O(N^2)
 *
 */
func InsertionSort(input []int) []int {
	result := make([]int, len(input))
	copy(result, input)

	for i := 1; i < len(result); i++ {
		temp_value := result[i]
		var position = i - 1

		for position >= 0 {
			if result[position] > temp_value {
				result[position+1] = result[position]
				position = position - 1
			} else {
				break
			}
		}
		result[position+1] = temp_value
	}

	return result
}
