package repository

import (
	"database/sql"
	"fmt"

	"github.com/MikhailMishutkin/Test_MediaSoft/internal/usecases"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1Vfcmrf1"
	dbname   = "testMS"
)

type Data struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBname   string `toml:"dbname"`
}

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

func (s *PersonRepository) Open() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.store.db = db

	return nil
}

func (s *GroupRepository) Open() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.store.db = db

	return nil
}

//Close
func (s *GroupRepository) Close() {

}

//Close
func (s *PersonRepository) Close() {

}
