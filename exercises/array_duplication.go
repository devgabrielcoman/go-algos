package exercises

func ArrayDuplication(array []int) []int {
	var result []int = []int{}

	var comparison map[int]int = make(map[int]int)
	for i := 0; i < len(array); i++ {
		item := array[i]
		_, exists := comparison[item]
		if exists {
			comparison[item] = comparison[item] + 1
		} else {
			comparison[item] = 1
		}
	}

	for key, value := range comparison {
		if value > 1 {
			result = append(result, key)
		}
	}

	return result
}
