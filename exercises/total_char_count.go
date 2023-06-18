package exercises

func TotalCharCount(input []string) int {
	if len(input) == 0 {
		return 0
	}

	if len(input) == 1 {
		return len(input[0])
	}
	return len(input[0]) + TotalCharCount(input[1:])
}
