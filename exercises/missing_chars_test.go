package exercises_test

import (
	"gabriel/algos/exercises"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MissingChars(t *testing.T) {

	t.Run("returns missing chars from text", func(t *testing.T) {
		input := "the quick brown fox jumps over the fence"
		result := exercises.MissingChars(input)
		assert.Equal(t, []string{"a", "d", "g", "l", "y", "z"}, result)
	})

	t.Run("returns no missing chars from string", func(t *testing.T) {
		input := "abcdefghijklmnopqrstuvwxyz"
		result := exercises.MissingChars(input)
		assert.Equal(t, []string{}, result)
	})
}
