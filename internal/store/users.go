package store

import "database/sql"

type UsersStore struct {
	db *sql.DB
}

func (u *UsersStore) Create() error {
	return nil
}
