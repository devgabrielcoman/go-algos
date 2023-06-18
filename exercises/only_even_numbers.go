package exercises

func OnlyEvenNumbrs(input []int) []int {
	if len(input) == 0 {
		return []int{}
	}

	var result = OnlyEvenNumbrs(input[1:])
	slice := result[:]

	if input[0]%2 == 0 {
		result = append([]int{input[0]}, slice...)
	}

	return result
}
