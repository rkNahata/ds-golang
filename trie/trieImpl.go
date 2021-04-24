package trie

import (
	"log"
)

type Node struct {
	Children  map[string]*Node
	endOfWord bool
}

func New() *Node {
	return &Node{
		Children:  make(map[string]*Node),
		endOfWord: false,
	}
}

func (trieNode *Node) InsertIterative(word string) {
	current := trieNode
	if current == nil {
		log.Fatal("root can't be nil")
		return
	}
	for i := range word {
		node := current.Children[string(word[i])]
		if node == nil {
			node = New()
			current.Children[string(word[i])] = node
		}
		current = node
	}
	current.endOfWord = true
}

func (trieNode *Node) Search(word string) bool {
	current := trieNode
	for i := range word {
		if current.Children[string(word[i])] == nil {
			return false
		}
		if current.Children[string(word[i])] != nil {
			current = current.Children[string(word[i])]
		}
	}
	return current.endOfWord
}

func (trieNode *Node) PrefixSearch(word string) (bool, *Node) {
	current := trieNode
	found := true
	for i := 0; i < len(word); i++ {
		if current.Children[string(word[i])] == nil {
			found = false
			break
		} else {
			current = current.Children[string(word[i])]
		}
	}
	return found, current
}

func (trieNode *Node) Autosuggest(prefix string) []string {
	var result []string
	if len(prefix) < 3 {
		return result
	}
	found, current := trieNode.PrefixSearch(prefix)
	if found {
		current.walkThePath(prefix, &result)
	}
	return result

}

func (current *Node) walkThePath(word string, result *[]string) {
	if current.endOfWord {
		*result = append(*result, word)
		return
	}
	for k, child := range current.Children {
		if child != nil {
			word += k
			child.walkThePath(word, result)
		}

	}

}
