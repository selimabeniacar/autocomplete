package dependencies

import (
	"autocomplete/suggestions/db"
	"database/sql"
)

type Dependencies interface {
	SuggestionsDB() *sql.DB
	TrieStore() db.TrieStore
	WordStore() db.WordStore
}

type dependencies struct {
	sDB *sql.DB
	ts  db.TrieStore
	ws  db.WordStore
}

func MakeDependencies() (Dependencies, error) {
	var d dependencies

	sDB, err := db.Connect()
	if err != nil {
		return nil, err
	}
	d.sDB = sDB

	d.ts = db.NewTrieStore(sDB)

	d.ws = db.NewwordStore(sDB)
	return &d, nil
}

func (d *dependencies) SuggestionsDB() *sql.DB {
	return d.sDB
}

func (d *dependencies) TrieStore() db.TrieStore {
	return d.ts
}

func (d *dependencies) WordStore() db.WordStore {
	return d.ws
}
