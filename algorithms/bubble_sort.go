package algorithms

/**
 * We continuously bubble up the largest value in the array by placing it on the nth, nth-1, nth-2 position successively.
 * We bubble up the value by constantly comparing two adjacent array values.
 *
 * It runs in:
 * 	best case - O(N).
 * 	avg case - O(N^2)
 * 	worst case - O(N^2)
 *
 * input: 			7 1 5 2 0
 * 1st run + 0: 1 7 5 2 0
 * 1st run + 1: 1 5 7 2 0
 * 1st run + 2: 1 5 2 7 0
 * 1st run + 3: 1 5 2 0 7
 * 2nd run:     1 2 0 5 7
 * 3nd run:			1 0 2 5 7
 * 4th run:			0 1 2 5 7 --> sorted
 */
func BubbleSort(input []int) []int {
	result := make([]int, len(input))
	copy(result, input)

	var sort_limit = len(result) - 1
	var sorted = false

	for !sorted {
		sorted = true
		for i := 0; i < sort_limit; i++ {
			current := result[i]
			next := result[i+1]
			if current > next {
				result[i+1] = current
				result[i] = next
				sorted = false
			}
		}
		sort_limit = sort_limit - 1
	}

	return result
}
