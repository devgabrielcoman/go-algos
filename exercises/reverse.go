package exercises

import "gabriel/algos/structures"

func Reverse(input string) string {
	var result = ""

	stack := structures.Stack[rune]()

	for _, char := range input {
		stack.Push(char)
	}

	for !stack.IsEmpty() {
		char := stack.Pop()
		if char != nil {
			result = result + string(*char)
		}
	}

	return result
}
