package algorithms

func QuickSort(input []int) []int {
	return inner_quicksort(input, 0, len(input)-1)
}

func inner_quicksort(input []int, left_pointer int, right_pointer int) []int {
	result := input[:]

	if right_pointer-left_pointer <= 0 {
		return result
	}

	pivot_index := partition(input, left_pointer, right_pointer)
	inner_quicksort(input, left_pointer, pivot_index-1)
	inner_quicksort(input, pivot_index+1, right_pointer)

	return result
}

func partition(input []int, left_pointer int, right_pointer int) int {
	pivot_index := right_pointer
	pivot := input[pivot_index]

	// we start the right pointer immediately to the left of the pivot
	right_pointer -= 1

	for {
		// move the left pointer to the right as long as it points to a value less than pivot
		for left_pointer <= len(input)-1 && input[left_pointer] < pivot {
			left_pointer += 1
		}

		// move the right pointer to the left as long as it points to a value greater than pivot
		for right_pointer >= 0 && input[right_pointer] > pivot {
			right_pointer -= 1
		}

		// we've stopped moving the pointers
		// if the left pointer has gone beyond the right pointer, we break
		if left_pointer >= right_pointer {
			break
		} else {
			// we swap the values of the left and right pointers
			tmp_left_pointer_val := input[left_pointer]
			input[left_pointer] = input[right_pointer]
			input[right_pointer] = tmp_left_pointer_val

			// we move the left pointer to the right
			left_pointer += 1
		}
	}

	// we swap the value of the left pointer with the pivot
	input[pivot_index] = input[left_pointer]
	input[left_pointer] = pivot

	return left_pointer
}
