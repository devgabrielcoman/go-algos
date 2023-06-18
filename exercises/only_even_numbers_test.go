package exercises_test

import (
	"gabriel/algos/exercises"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OnlyEvenNumbers(t *testing.T) {

	t.Run("returns empty array for empty input", func(t *testing.T) {
		input := []int{}
		result := exercises.OnlyEvenNumbrs(input)
		assert.Equal(t, []int{}, result)
	})

	t.Run("returns only even numbers", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 10}
		result := exercises.OnlyEvenNumbrs(input)
		assert.Equal(t, []int{2, 4, 6, 8, 10}, result)
	})

	t.Run("returns empty for an array of odd numbers", func(t *testing.T) {
		input := []int{1, 3, 5, 7, 13, 21}
		result := exercises.OnlyEvenNumbrs(input)
		assert.Equal(t, []int{}, result)
	})
}
