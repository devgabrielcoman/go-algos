package exercises

import "strings"

func FirstNonDuplicate(input string) string {
	inputArray := strings.Split(input, "")
	var inputMap map[string]int = make(map[string]int)

	for _, char := range inputArray {
		_, exists := inputMap[char]
		if exists {
			inputMap[char] = inputMap[char] + 1
		} else {
			inputMap[char] = 1
		}
	}

	for _, char := range inputArray {
		if inputMap[char] == 1 {
			return char
		}
	}

	return ""
}
