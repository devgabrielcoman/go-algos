package exercises

/**
 * Calculating the factorial of a "number" means doing:
 * 1 * 2 * ... * (number - 1) * number
 */
func Factorial(number int) int {
	if number <= 1 {
		return number
	} else {
		return number * Factorial(number-1)
	}
}
