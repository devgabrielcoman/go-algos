package structures_test

import (
	"fmt"
	"gabriel/algos/structures"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func randomStrings(number_of_strings, min_len, max_len int) []string {
	rand.Seed(time.Now().UnixNano())

	strings := make([]string, number_of_strings)
	for i := 0; i < number_of_strings; i++ {
		length := rand.Intn(max_len-min_len+1) + min_len
		str := make([]byte, length)
		for j := 0; j < length; j++ {
			str[j] = byte(rand.Intn(26) + 97) // Generate lowercase ASCII characters
		}
		strings[i] = string(str)
	}

	return strings
}

func Test_HashMap(t *testing.T) {
	t.Run("can create a new hash map", func(t *testing.T) {
		hashmap := structures.HashMap[int]()
		assert.NotNil(t, hashmap)
	})

	t.Run("can hash any string into a value from 1 to 9", func(t *testing.T) {
		hashmap := structures.HashMap[int]()
		input := randomStrings(100, 5, 25)

		for _, value := range input {
			hash := hashmap.Hash(value)
			fmt.Println(hash)
			assert.True(t, hash >= 1)
			assert.True(t, hash <= 9)
		}
	})

	t.Run("can set and read from the hash map", func(t *testing.T) {
		hashmap := structures.HashMap[int]()

		res1 := hashmap.Get("key")
		assert.Nil(t, res1)

		hashmap.Set("key", 5)
		res2 := hashmap.Get("key")
		assert.Equal(t, 5, *res2)

		hashmap.Set("key-2", 7)
		hashmap.Set("key-3", 10)
		res3 := hashmap.Get("key-2")
		res4 := hashmap.Get("key-3")
		res5 := hashmap.Get("unknown-key")

		assert.Equal(t, 7, *res3)
		assert.Equal(t, 10, *res4)
		assert.Nil(t, res5)
	})

	t.Run("can override value with same key", func(t *testing.T) {
		hashmap := structures.HashMap[int]()

		hashmap.Set("key", 5)
		hashmap.Set("key", 7)

		res := hashmap.Get("key")
		assert.Equal(t, 7, *res)
	})
}
