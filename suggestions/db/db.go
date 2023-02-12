package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	cfg := mysql.Config{
		User:   "root",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "suggestions",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

type FreqData struct {
	Word string
	Freq int
}

type TrieStore interface {
	GetTrie() ([]byte, error)
	PutTrie([]byte) error
	SeedTrie([]byte) error
}

type WordStore interface {
	GetEntries() ([]FreqData, error)
	PutEntry(string) error
}

func NewTrieStore(db *sql.DB) TrieStore {
	return &trieStore{db: db}
}

func NewwordStore(db *sql.DB) WordStore {
	return &wordStore{db: db}
}

type trieStore struct {
	db *sql.DB
}

func (ts *trieStore) GetTrie() ([]byte, error) {
	var b []byte
	row := ts.db.QueryRow("SELECT trie FROM trie WHERE ID=?", 1)
	if err := row.Scan(&b); err != nil {
		if err == sql.ErrNoRows {
			return []byte{}, nil
		}
		return nil, err
	}
	return b, nil
}

func (ts *trieStore) PutTrie(b []byte) error {
	_, err := ts.db.Exec("UPDATE trie SET trie =? WHERE ID=?", b, 1)
	if err != nil {
		return err
	}
	return nil
}

func (ts *trieStore) SeedTrie(t []byte) error {
	// First time inserting the trie
	_, err := ts.db.Exec("INSERT INTO trie (trie) VALUES (?)", t)
	if err != nil {
		return err
	}
	return nil
}

type wordStore struct {
	db *sql.DB
}

func (ss *wordStore) GetEntries() ([]FreqData, error) {
	rows, err := ss.db.Query("SELECT Word, Freq FROM words")
	if err != nil {
		return []FreqData{}, err
	}

	var res []FreqData
	for rows.Next() {
		var fq FreqData
		if err := rows.Scan(&fq.Word, &fq.Freq); err != nil {
			return []FreqData{}, err
		}
		res = append(res, fq)
	}
	return res, nil
}

func (ss *wordStore) PutEntry(word string) error {
	res, err := ss.db.Exec("UPDATE words SET Freq = Freq+1 WHERE Word=?", word)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	// First time inserting the word
	if affected == 0 {
		_, err = ss.db.Exec("INSERT INTO words (Word, Freq) VALUES (?, ?)", word, 1)
		if err != nil {
			return err
		}
	}
	return nil
}
