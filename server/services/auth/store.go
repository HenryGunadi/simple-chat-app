package auth

import "database/sql"

type Store struct {
	store *sql.DB
}

func NewAuthStore(store *sql.DB) *Store {
	return &Store{store: store}
}
