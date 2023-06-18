package exercises_test

import (
	"gabriel/algos/exercises"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Revers(t *testing.T) {

	t.Run("returns empty for empty input", func(t *testing.T) {
		input := ""
		result := exercises.Reverse(input)
		assert.Equal(t, "", result)
	})

	t.Run("returns a reversed string", func(t *testing.T) {
		input := "abcdef"
		result := exercises.Reverse(input)
		assert.Equal(t, "fedcba", result)
	})
}
