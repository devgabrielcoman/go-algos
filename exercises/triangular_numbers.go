package exercises

func TriangularNumbers(number int) int {
	if number < 0 {
		return 0
	}

	switch number {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 3
	case 3:
		return 6
	case 4:
		return 10
	case 5:
		return 15
	case 6:
		return 21
	default:
		return number + TriangularNumbers(number-1)
	}
}
