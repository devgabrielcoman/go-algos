package exercises

import (
	"sort"
	"strings"
)

func MissingChars(input string) []string {
	var result []string = []string{}

	alphabet := strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	var alphabetMap map[string]bool = make(map[string]bool)

	for _, char := range alphabet {
		alphabetMap[char] = true
	}

	chars := strings.Split(input, "")
	var charMap map[string]bool = make(map[string]bool)

	for _, char := range chars {
		charMap[char] = true
	}

	for key := range alphabetMap {
		_, exists := charMap[key]
		if !exists {
			result = append(result, key)
		}
	}

	sort.Strings(result)
	return result
}
