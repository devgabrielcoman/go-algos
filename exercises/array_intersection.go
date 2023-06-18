package exercises

/**
 * Efficiently calculate the intersection of two arrays of different sizes in O(N)
 */
func ArrayIntersection(array1 []int, array2 []int) []int {
	var result []int = []int{}

	var bigArray []int
	var smallArray []int
	if len(array1) > len(array2) {
		bigArray = array1
		smallArray = array2
	} else {
		bigArray = array2
		smallArray = array1
	}

	var comparator map[int]bool = make(map[int]bool)
	for i := 0; i < len(bigArray); i++ {
		comparator[bigArray[i]] = true
	}

	for j := 0; j < len(smallArray); j++ {
		target := smallArray[j]
		_, exists := comparator[target]
		if exists {
			result = append(result, target)
		}
	}

	return result
}
