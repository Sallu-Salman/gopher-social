package store

import "database/sql"

type Storage struct {
	Posts PostsStore
	Users UsersStore
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: PostsStore{db: db},
		Users: UsersStore{db: db},
	}
}
