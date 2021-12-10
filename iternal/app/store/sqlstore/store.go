package sqlstore

import (
	"database/sql" //...
	"http-rest-api/iternal/app/store"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/lib/pq"
)

//Store...
type Store struct {
	db             *sql.DB
	UserRepository *UserRepository
}

//New...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

//User...
func (s *Store) User() store.UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}

	s.UserRepository = &UserRepository{
		store: s,
	}

	return s.UserRepository
}
