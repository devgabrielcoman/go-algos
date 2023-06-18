package exercises

/**
 * Calculates the Nth fibonacci number.
 * This is an elegant but slow approach, since performance is O(2^N) because of the two recursive calls
 */
func FibonacciRaw(number int) int {
	if number <= 0 {
		return 0
	}
	if number == 1 {
		return 1
	}
	return FibonacciRaw(number-1) + FibonacciRaw(number-2)
}

/**
 * Calculates the Nth fibonacci number.
 * This is an improved approach, using DynamicProgramming "memoisation".
 * This runs in 2N - 1, which is O(N)
 */
func FibonacciDynamic(number int, memo map[int]int) int {
	if number <= 0 {
		return 0
	}
	if number == 1 {
		return 1
	}

	_, exists := memo[number]

	if !exists {
		memo[number] = FibonacciDynamic(number-1, memo) + FibonacciDynamic(number-2, memo)
	}

	return memo[number]
}
