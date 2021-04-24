package trie

import "testing"

func TestInsertIterative(t *testing.T) {
	words := []string{"tab", "bullet", "boost", "boom", "i"}
	table := []struct {
		searchWord string
		isPresent  bool
	}{
		{"tab", true},
		{"bullet", true},
		{"boost", true},
		{"boom", true},
		{"i", true},
		{"ing", false},
		{"bull", true},
	}
	trie := New()
	for _, w := range words {
		trie.InsertIterative(w)
	}

	for _, w := range table {
		actual := trie.Search(w.searchWord)
		if actual != w.isPresent {
			t.Errorf("expected %s to be present but it is not!!!", w.searchWord)
		}
	}
}

func Test_AutoSuggest(t *testing.T) {
	words := []string{"tab12", "tab123", "tab23", "tab1", "bull", "i"}
	table := []struct {
		searchWord string
		count      int
	}{
		{"tab", 4},
		{"", 0},
		{"bull", 1},
	}
	trie := New()
	for _, w := range words {
		trie.InsertIterative(w)
	}

	for _, w := range table {
		result := trie.Autosuggest(w.searchWord)
		if len(result) != w.count {
			t.Errorf("expected result to be %d but got %d!!!", w.count, len(result))
		}
	}
}
