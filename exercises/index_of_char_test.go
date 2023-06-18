package exercises_test

import (
	"gabriel/algos/exercises"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IndexOfChar(t *testing.T) {
	t.Run("returns -1 for an empty input", func(t *testing.T) {
		input := ""
		result := exercises.IndexOfChar(input, 'x', 0)
		assert.Equal(t, -1, result)
	})

	t.Run("returns -1 for a one char string that does not contain the target char", func(t *testing.T) {
		input := "a"
		result := exercises.IndexOfChar(input, 'x', 0)
		assert.Equal(t, -1, result)
	})

	t.Run("returns -1 for a multi char string that does not contain the target char", func(t *testing.T) {
		input := "abcdef"
		result := exercises.IndexOfChar(input, 'x', 0)
		assert.Equal(t, -1, result)
	})

	t.Run("returns the first position of the target char in the input string", func(t *testing.T) {
		input := "abcxefx1x"
		result := exercises.IndexOfChar(input, 'x', 0)
		assert.Equal(t, 3, result)
	})

	t.Run("returns the first position of the target char in case it is the last one", func(t *testing.T) {
		input := "abcuefi1x"
		result := exercises.IndexOfChar(input, 'x', 0)
		assert.Equal(t, 8, result)
	})
}
