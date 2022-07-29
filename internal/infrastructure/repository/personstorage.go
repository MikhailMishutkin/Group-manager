package repository

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/MikhailMishutkin/Test_MediaSoft/internal/domain"
)

type PersonRepository struct {
	store  *Store
	logger *logrus.Logger
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

// list all person in database
func (r *PersonRepository) GetList() (jsonData []byte, err error) {
	r.logger = logrus.New()
	q := `
	SELECT json_agg(s) FROM (
		SELECT id, person_name, surname, year_of_birth, groupname
		from persons 
	) s;`
	if err := r.store.db.QueryRow(q).Scan(&jsonData); err != nil {
		r.logger.Printf("Error GetList of persons: %s", err)
		return nil, err
	}
	r.logger.Tracef("Try to get list of persons", q)
	return jsonData, nil
}

// update person's group
func (r *PersonRepository) UpdatePerson(id int, gn string) {
	q := `UPDATE persons SET groupname = $2 WHERE id = $1`

	_, err := r.store.db.Exec(q, id, gn)
	if err != nil {
		log.Fatalf("error to update person: %s", err)
	}
}

func (r *PersonRepository) GetListAll() {

}

// delete person from database
func (r *PersonRepository) DeletePerson(id int) {
	q := `DELETE FROM persons where id = $1`

	_, err := r.store.db.Exec(q, id)
	if err != nil {
		log.Fatalf("error to delete person: %s", err)
	}
}
