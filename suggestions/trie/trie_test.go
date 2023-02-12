package trie_test

import (
	"autocomplete/suggestions/db"
	"autocomplete/suggestions/trie"
	"autocomplete/suggestions/triepb"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTopNSuggestions(t *testing.T) {
	tests := map[string]struct {
		fd       []db.FreqData
		n        int
		expected []string
	}{
		"empty trie": {
			fd:       []db.FreqData{},
			n:        5,
			expected: []string{},
		},
		"all starting with be": {
			fd: []db.FreqData{
				{Word: "be", Freq: 2},
				{Word: "bear", Freq: 5},
				{Word: "beer", Freq: 3},
				{Word: "beast", Freq: 6},
				{Word: "beard", Freq: 1},
			},
			n:        3,
			expected: []string{"beast", "bear", "beer"},
		},
	}

	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			root := &triepb.Trie{Fanout: make([]*triepb.Trie, trie.AlphabetLength), Freqs: 0, End: false}
			for _, fd := range v.fd {
				trie.Add(root, fd.Word, fd.Freq)
			}
			topN, err := trie.GetTopNSuggestionsHelper(root, "", v.n)
			assert.Nil(t, err)
			assert.True(t, len(topN) == len(v.expected))

			s := []string{}
			for _, word := range topN {
				s = append(s, word.Word)
			}
			assert.Equal(t, s, v.expected)
		})
	}
}
