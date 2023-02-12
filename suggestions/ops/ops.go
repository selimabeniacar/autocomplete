package ops

import (
	"autocomplete/suggestions/db"
	"autocomplete/suggestions/dependencies"
	"autocomplete/suggestions/trie"
	"log"
	"time"
)

func GetTopNSuggestions(ts db.TrieStore, word string, n int) ([]string, error) {
	return trie.GetTopNSuggestions(ts, word, n)
}

func AddWord(ws db.WordStore, word string) error {
	// Validity check
	for _, v := range word {
		if !(v >= 'a' && v <= 'z') {
			return nil
		}
	}
	return ws.PutEntry(word)
}

func BuildTrieForever(d dependencies.Dependencies) {
	for {
		// Get the words and their frequencies
		entries, err := d.WordStore().GetEntries()
		if err != nil {
			log.Println("Error getting word data")
			time.Sleep(10 * time.Second)
			continue
		}

		t := trie.NewTrie()
		for _, v := range entries {
			trie.Add(t, v.Word, v.Freq)
		}

		// Serialize the trie and persist to a database
		b, err := trie.SerializeTrie(t)
		if err != nil {
			log.Println("Error serializing the trie")
			time.Sleep(10 * time.Second)
			continue
		}

		err = d.TrieStore().PutTrie(b)
		if err != nil {
			log.Println("Error persisting the trie")
			time.Sleep(10 * time.Second)
			continue
		}

		time.Sleep(5 * time.Second)
	}
}

func SeedTrie(d dependencies.Dependencies) error {
	b, err := d.TrieStore().GetTrie()
	if err != nil {
		return err
	}
	// Seed if there's no trie persisted
	if len(b) == 0 {
		t := trie.NewTrie()
		b, err = trie.SerializeTrie(t)
		if err != nil {
			return err
		}
		err = d.TrieStore().SeedTrie(b)
		if err != nil {
			return err
		}
	}
	return nil
}
