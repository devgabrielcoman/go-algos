package exercises_test

import (
	"gabriel/algos/exercises"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TotalCharCount(t *testing.T) {

	t.Run("returns 0 for empty input", func(t *testing.T) {
		input := []string{}
		result := exercises.TotalCharCount(input)
		assert.Equal(t, 0, result)
	})

	t.Run("returns 0 for array of empty strings", func(t *testing.T) {
		input := []string{"", "", ""}
		result := exercises.TotalCharCount(input)
		assert.Equal(t, 0, result)
	})

	t.Run("returns correct length for valid input", func(t *testing.T) {
		input := []string{"ab", "c", "def", "ghij"}
		result := exercises.TotalCharCount(input)
		assert.Equal(t, 10, result)
	})
}
