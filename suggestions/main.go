package main

import (
	"autocomplete/suggestions/dependencies"
	"autocomplete/suggestions/handlers"
	"autocomplete/suggestions/ops"
	"log"
	"net/http"
)

func main() {
	d, err := dependencies.MakeDependencies()
	if err != nil {
		log.Fatal(err)
	}

	err = ops.SeedTrie(d)
	if err != nil {
		log.Fatal(err)
	}

	go ops.BuildTrieForever(d)

	http.HandleFunc("/", handlers.MainPageHandler)
	http.HandleFunc("/search", handlers.AddWordHandler(d.WordStore()))
	http.HandleFunc("/gettopn", handlers.TopNSuggestionHandler(d.TrieStore()))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
