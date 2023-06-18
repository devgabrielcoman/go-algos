package exercises_test

import (
	"gabriel/algos/exercises"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TriangularNumbers(t *testing.T) {

	t.Run("returns 45 for the 8th number", func(t *testing.T) {
		result := exercises.TriangularNumbers(9)
		assert.Equal(t, 45, result)
	})

	t.Run("returns 36 for the 8th number", func(t *testing.T) {
		result := exercises.TriangularNumbers(8)
		assert.Equal(t, 36, result)
	})

	t.Run("returns 28 for the 7th number", func(t *testing.T) {
		result := exercises.TriangularNumbers(7)
		assert.Equal(t, 28, result)
	})

	t.Run("returns 21 for the 6th number", func(t *testing.T) {
		result := exercises.TriangularNumbers(6)
		assert.Equal(t, 21, result)
	})

	t.Run("returns 15 for the 5th number", func(t *testing.T) {
		result := exercises.TriangularNumbers(5)
		assert.Equal(t, 15, result)
	})

	t.Run("returns 10 for the 4th number", func(t *testing.T) {
		result := exercises.TriangularNumbers(4)
		assert.Equal(t, 10, result)
	})

	t.Run("returns 6 for the 3rd number", func(t *testing.T) {
		result := exercises.TriangularNumbers(3)
		assert.Equal(t, 6, result)
	})

	t.Run("returns 3 for the 2nd number", func(t *testing.T) {
		result := exercises.TriangularNumbers(2)
		assert.Equal(t, 3, result)
	})

	t.Run("returns 1 for the 1st number", func(t *testing.T) {
		result := exercises.TriangularNumbers(1)
		assert.Equal(t, 1, result)
	})
}
