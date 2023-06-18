package exercises_test

import (
	"gabriel/algos/exercises"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FirstNonDuplicate(t *testing.T) {
	t.Run("returns empty for an empty input", func(t *testing.T) {
		input := ""
		result := exercises.FirstNonDuplicate(input)
		assert.Equal(t, "", result)
	})

	t.Run("returns empty for an input without duplication", func(t *testing.T) {
		input := "aazzbb"
		result := exercises.FirstNonDuplicate(input)
		assert.Equal(t, "", result)
	})

	t.Run("returns first non duplicate letter in the order it appears in the string", func(t *testing.T) {
		input := "wowzza"
		result := exercises.FirstNonDuplicate(input)
		assert.Equal(t, "o", result)
	})
}
