package repository

import (
	"database/sql"

	"github.com/MikhailMishutkin/Test_MediaSoft/internal/usecases"
	_ "github.com/lib/pq"
)

type Store struct {
	db               *sql.DB
	personRepository *PersonRepository
	groupRepository  *GroupRepository
}

type Storer interface {
	Person() usecases.PersonRepository
	Group() usecases.GroupRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Person() usecases.PersonRepository {
	if s.personRepository == nil {
		s.personRepository = &PersonRepository{
			store: s,
		}
		return s.personRepository
	}

	return s.personRepository
}

func (s *Store) Group() usecases.GroupRepository {
	if s.groupRepository == nil {
		s.groupRepository = &GroupRepository{
			store: s,
		}
		return s.groupRepository
	}

	return s.groupRepository
}
