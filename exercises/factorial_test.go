package exercises_test

import (
	"gabriel/algos/exercises"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Factorial(t *testing.T) {
	t.Run("returns the number if it's negative", func(t *testing.T) {
		result := exercises.Factorial(-1)
		assert.Equal(t, -1, result)
	})

	t.Run("returns the number if it's 1", func(t *testing.T) {
		result := exercises.Factorial(1)
		assert.Equal(t, 1, result)
	})

	t.Run("returns the factorial of a positive number", func(t *testing.T) {
		result := exercises.Factorial(3)
		assert.Equal(t, 6, result)
	})
}
