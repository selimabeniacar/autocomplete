package trie

import (
	"autocomplete/suggestions/db"
	"autocomplete/suggestions/triepb"

	"google.golang.org/protobuf/proto"
)

// Only supports lowercase English letters for now
const AlphabetLength = 26

// TODO(aben): Implement own Trie type so protobuf details are not leaked to this API
func NewTrie() *triepb.Trie {
	return &triepb.Trie{Fanout: make([]*triepb.Trie, AlphabetLength), Freqs: 0, End: false}
}

// Add adds the word with given frequency to Trie
func Add(t *triepb.Trie, word string, freq int) {
	if len(word) == 0 {
		return
	}

	c := word[0]
	if t.GetFanout()[c-'a'] == nil {
		t.GetFanout()[c-'a'] = NewTrie()
	}
	if len(word) == 1 {
		t.GetFanout()[c-'a'].End = true
		t.GetFanout()[c-'a'].Freqs = int32(freq)
		return
	}
	Add(t.GetFanout()[c-'a'], word[1:], freq)
}

// Deserialize to triepb.Trie from protobuf
func DeserializeTrie(encoded []byte) (*triepb.Trie, error) {
	t := &triepb.Trie{}
	err := proto.Unmarshal(encoded, t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// Serialize to protobuf from triepb.Trie
func SerializeTrie(t *triepb.Trie) ([]byte, error) {
	b, err := proto.Marshal(t)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GetTopNSuggestions(ts db.TrieStore, word string, n int) ([]string, error) {
	// Validity check
	for _, v := range word {
		if !(v >= 'a' && v <= 'z') {
			return []string{}, nil
		}
	}

	tb, err := ts.GetTrie()
	if err != nil {
		return nil, err
	}
	t, err := DeserializeTrie(tb)
	if err != nil {
		return nil, err
	}

	wc := word
	curr := t
	// Can't check for nil because repeated nil messages are serialized and deserialized as empty messages.
	// That's why we are checking the fanout length here(0 for empty message / nil)
	for len(curr.GetFanout()) > 0 && len(word) > 0 {
		c := word[0]
		curr = curr.GetFanout()[c-'a']
		word = word[1:]
	}

	if len(word) > 0 {
		return []string{}, nil
	}

	mostNFreq, err := GetTopNSuggestionsHelper(curr, wc, n)
	if err != nil {
		return nil, err
	}

	var res []string
	for _, v := range mostNFreq {
		res = append(res, v.Word)
	}
	return res, nil
}

func GetTopNSuggestionsHelper(root *triepb.Trie, word string, n int) ([]db.FreqData, error) {
	if root == nil {
		return []db.FreqData{}, nil
	}

	indices := make([]int, AlphabetLength)
	children := make([][]db.FreqData, AlphabetLength)

	var res []db.FreqData
	node := root
	for i, v := range node.GetFanout() {
		v := v
		var err error
		children[i], err = GetTopNSuggestionsHelper(v, word+string('a'+byte(i)), n)
		if err != nil {
			return nil, err
		}
	}

	// O(26 * n), n is small so should be fine.
	for len(res) < n {
		var mostFreq db.FreqData
		var index int
		for i, v := range indices {
			if len(children[i]) > v && children[i][v].Freq > mostFreq.Freq {
				mostFreq = children[i][v]
				index = i
			}
		}
		if mostFreq.Freq == 0 {
			break
		}
		res = append(res, mostFreq)
		indices[index]++
	}

	if node.End {
		l := len(res)
		if l == 0 {
			res = append(res, db.FreqData{Word: word, Freq: int(node.GetFreqs())})
		} else {
			for i, v := range res {
				if node.GetFreqs() > int32(v.Freq) {
					right := res[i:l:l]
					res = append(res[:i], append([]db.FreqData{{Word: word, Freq: int(node.GetFreqs())}}, res[i:]...)...)
					res = append(res, right...)
					break
				}
			}
		}
		// If we couldn't insert it and there's room insert to end.
		if l == len(res) && l < n {
			res = append(res, db.FreqData{Word: word, Freq: int(node.GetFreqs())})
		}
	}
	return res, nil
}
