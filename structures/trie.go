package structures

type trieNode struct {
	children map[rune]trieNode
}

func TrieNode() *trieNode {
	return &trieNode{
		children: make(map[rune]trieNode),
	}
}

type trie struct {
	root *trieNode
}

func Trie() *trie {
	return &trie{
		root: nil,
	}
}

func (t *trie) IsEmpty() bool {
	return t.root == nil
}

func (t *trie) Insert(word string) {
	if t.IsEmpty() {
		t.root = TrieNode()
	}

	var current_node = t.root

	terminator := "*"
	full_word := word + terminator

	for _, character := range full_word {
		node, exists := current_node.children[character]

		if !exists {
			next_node := TrieNode()
			current_node.children[character] = *next_node
			current_node = next_node
		} else {
			current_node = &node
		}
	}
}

func (t *trie) Autocomplete(prefix string) []string {
	if t.IsEmpty() {
		return []string{}
	}

	node := t.search(prefix)

	if node == nil {
		return []string{}
	}

	words := []string{}
	t.collect(node, "", &words)

	var result = []string{}
	for _, word := range words {
		result = append(result, prefix+word)
	}
	return result
}

func (t *trie) Collect() []string {
	if t.IsEmpty() {
		return []string{}
	}

	words := []string{}
	t.collect(t.root, "", &words)
	return words
}

func (t *trie) collect(current_node *trieNode, word string, words *[]string) {
	for character, node := range current_node.children {
		if character == '*' {
			*words = append(*words, word)
		} else {
			t.collect(&node, word+string(character), words)
		}
	}
}

func (t *trie) search(prefix string) *trieNode {
	var current_node = t.root

	for _, char := range prefix {
		new_node, exists := current_node.children[char]
		if exists {
			current_node = &new_node
		} else {
			return nil
		}
	}

	return current_node
}
