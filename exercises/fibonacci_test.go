package exercises_test

import (
	"gabriel/algos/exercises"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Fibonacci(t *testing.T) {
	t.Run("returns 0 if it's negative", func(t *testing.T) {
		result := exercises.FibonacciRaw(-1)
		assert.Equal(t, 0, result)
	})

	t.Run("returns the correct fibonacci numbers for unoptimised algorithm", func(t *testing.T) {
		numbers := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
		for i, val := range numbers {
			result := exercises.FibonacciRaw(i)
			assert.Equal(t, val, result)
		}
	})

	t.Run("returns the correct fibonacci numbers for dynamic algorithm", func(t *testing.T) {
		numbers := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
		for i, val := range numbers {
			result := exercises.FibonacciDynamic(i, make(map[int]int))
			assert.Equal(t, val, result)
		}
	})
}
