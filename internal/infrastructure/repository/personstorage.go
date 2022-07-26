package repository

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/MikhailMishutkin/Test_MediaSoft/internal/domain"
)

type PersonRepository struct {
	store *Store
}

func NewPersonRepository(store *Store) *PersonRepository {
	return &PersonRepository{
		store: store,
	}
}

// Create Person

func (r *PersonRepository) MakePerson(p *domain.Person) (*domain.Person, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO persons (person_name, surname, year_of_birth, groupname) VALUES ($1, $2, $3, $4) RETURNING id",
		p.Name,
		p.Surname,
		p.YearOfBirth,
		p.GroupName,
	).Scan(&p.ID); err != nil {
		return nil, err
	}

	return p, nil

}

func (r *PersonRepository) GetList() (jsonData []byte, err error) {
	logger := logrus.New()
	logger.Trace("Try to get list of persons")

	q := `
	SELECT json_agg(s) FROM (
		SELECT id, person_name, surname, year_of_birth, groupname
		from persons 
	) s;`
	if err := r.store.db.QueryRow(q).Scan(&jsonData); err != nil {
		logger.Printf("Error GetList of persons: %s", err)
		return nil, err
	}

	return jsonData, nil
}

func (r *PersonRepository) GetListAll() {

}

func (r *PersonRepository) DeletePerson() {

}

func (r *PersonRepository) UpdatePerson() {

}
