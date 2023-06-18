package structures_test

import (
	"gabriel/algos/structures"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Trie(t *testing.T) {
	t.Run("can create a trie", func(t *testing.T) {
		trie := structures.Trie()
		assert.NotNil(t, trie)
	})

	t.Run("can insert a word into the trie", func(t *testing.T) {
		trie := structures.Trie()

		trie.Insert("word")
		trie.Insert("wordless")
		result1 := trie.Collect()
		assert.True(t, containsValue(result1, "word"))
		assert.True(t, containsValue(result1, "wordless"))

		trie.Insert("ape")
		trie.Insert("apple")
		result2 := trie.Collect()
		assert.True(t, containsValue(result2, "apple"))
		assert.True(t, containsValue(result2, "ape"))
		assert.True(t, containsValue(result2, "word"))
		assert.True(t, containsValue(result2, "wordless"))
	})

	t.Run("can use autocomplete", func(t *testing.T) {
		trie := structures.Trie()

		trie.Insert("word")
		trie.Insert("wordless")

		result1 := trie.Autocomplete("wo")
		assert.True(t, containsValue(result1, "wordless"))
		assert.True(t, containsValue(result1, "word"))

		result2 := trie.Autocomplete("ape")
		assert.Equal(t, []string{}, result2)

		trie.Insert("ape")
		trie.Insert("apple")

		result3 := trie.Autocomplete("ap")
		assert.True(t, containsValue(result3, "apple"))
		assert.True(t, containsValue(result3, "ape"))
	})
}

func containsValue(arr []string, value string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}
