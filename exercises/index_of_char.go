package exercises

func IndexOfChar(input string, target rune, start int) int {
	if len(input) == 0 {
		return -1
	}

	if input[0] == byte(target) {
		return start
	}

	return IndexOfChar(input[1:], target, start+1)
}
