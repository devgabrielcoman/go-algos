package algorithms

/**
 * We're continously splitting a sorted array into two halves, based on the midpoint;
 * if the searched value is greater than the value at midpoint => we search in the upper half
 * if the searched value is less than the value at midpoint => we search in the lower half
 * we do this until we either find the desired value or not.
 * It runs in O(logN).
 *
 * input:		1 2 3 5 7 13 21
 * search: 	13
 * 1st run: 7 13 21
 * 2nd run: 13 --> found
 */
func BinarySearch(input []int, search int) bool {
	if len(input) == 0 {
		return false
	}

	if len(input) == 1 {
		return input[0] == search
	}

	var start = 0
	var end = len(input) - 1

	for start <= end {
		mindpoint := (end + start) / 2
		if search == input[mindpoint] {
			return true
		} else if search > input[mindpoint] {
			start = mindpoint + 1
		} else {
			end = mindpoint - 1
		}
	}

	return false
}
