package handlers

import (
	"autocomplete/suggestions/db"
	"autocomplete/suggestions/ops"
	"encoding/json"
	"log"
	"net/http"
)

func TopNSuggestionHandler(ts db.TrieStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		word := r.FormValue("search")
		suggs, err := ops.GetTopNSuggestions(ts, word, 5)
		if err != nil {
			log.Fatal(err)
		}
		b, err := json.Marshal(suggs)
		if err != nil {
			log.Fatal(err)
		}
		//w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}

func AddWordHandler(ws db.WordStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		word := r.FormValue("search")
		err := ops.AddWord(ws, word)
		if err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./ui/templates/static/index.html")
}
