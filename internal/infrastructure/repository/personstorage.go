package repository

import (
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/MikhailMishutkin/Test_MediaSoft/internal/domain"
)

//структура для вывода данных пользователю
type listPerson struct {
	Group     string           `json:"group"`
	Persons   []*domain.Person `json:"persons"`
	Subgroups []subgroups      `json:"subgroups"`
}

// подструктура для вывода данных пользователю
type subgroups struct {
	Group   string           `json:"group"`
	Persons []*domain.Person `json:"persons"`
}

// структура с инъекцией хранилища
type PersonRepository struct {
	store  *Store
	logger *logrus.Logger
}

// конструктор
func NewPersonRepository(store *Store) *PersonRepository {
	return &PersonRepository{
		store: store,
	}
}

// создаём запись о человеке
func (r *PersonRepository) CreatePerson(p *domain.Person) (*domain.Person, error) {
	r.logger = logrus.New()
	var gs []string
	var s string

	//извлекаем значения существующих групп
	rows, err := r.store.db.Query("SELECT groupname FROM groups")
	if err != nil {
		r.logger.Printf("Error to get groupnames from db: %s", err)
		return nil, err
	}
	defer rows.Close()

	//складываем группы в массив
	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			return p, err
		}
		gs = append(gs, s)
	}

	// проверяем существование группы в бд
	for _, v := range gs {
		if p.GroupName == v {

			// если группа существует, то добавляем пользователя
			if err := r.store.db.QueryRow(
				"INSERT INTO persons (person_name, surname, year_of_birth, groupname) VALUES ($1, $2, $3, $4) RETURNING id",
				p.Name,
				p.Surname,
				p.YearOfBirth,
				p.GroupName,
			).Scan(&p.ID); err != nil {
				return nil, err
			}
			// добавляем человека в таблицу группы
			r.store.db.QueryRow(
				`UPDATE groups 
				SET members = members + 1
				WHERE groupname = $1`, p.GroupName)
			return p, nil
		} else {
			continue
		}
	}
	err = fmt.Errorf("no such group, try again")

	return p, err
}

// обновляем запись о человеке, можно обновить только группу
func (r *PersonRepository) UpdatePerson(id int, gn string) error {
	r.logger = logrus.New()
	err := r.store.db.QueryRow("SELECT groupname FROM persons WHERE id = $1", id).Scan(&gn)
	if err != nil {
		r.logger.Printf("Error to get groupnames from db: %s", err)
		return err
	}
	q := `UPDATE persons SET groupname = $2 WHERE id = $1`

	_, err = r.store.db.Exec(q, id, gn)
	if err != nil {
		log.Fatalf("error to update person: %s", err)
	}

	return nil
}

// удаляем запись о человеке по id
func (r *PersonRepository) DeletePerson(id int) error {
	r.logger = logrus.New()
	// используем queryrow чтобы просканировать группу в переменную
	var gn string
	err := r.store.db.QueryRow("SELECT groupname FROM persons WHERE id = $1", id).Scan(&gn)
	if err != nil {
		r.logger.Printf("Error to get groupnames from db: %s", err)
		return err
	}

	_, err = r.store.db.Exec(`DELETE FROM persons where id = $1`, id)
	if err != nil {
		r.logger.Info("error to delete person: %s", err)
		return err
	}
	// уменьшаем счётчик  людей в группе
	r.store.db.QueryRow(
		`UPDATE groups 
		SET members = members - 1
		WHERE groupname = $1`, gn)

	return nil
}

// отображает список людей в определённой группе только привязанных к данной
// и список людей со всеми дочерними группами
func (r *PersonRepository) GetList(gn string) (jsonData []byte, err error) {
	r.logger = logrus.New()
	var l listPerson
	var sub subgroups
	var p *domain.Person
	var n, sur string
	var id, y int

	q := "SELECT id, person_name, surname, year_of_birth FROM persons WHERE groupname = $1"
	//извлекаем данные о людях в группе для группировки в структуру
	rows, err := r.store.db.Query(q, gn)
	if err != nil {
		r.logger.Printf("Error to get person data from db: %s", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&id, &n, &sur, &y); err != nil {
			r.logger.Printf("trouble with rows.Next: %s", err)
			return nil, err
		}
		p = &domain.Person{
			ID:          id,
			Name:        n,
			Surname:     sur,
			YearOfBirth: y,
			GroupName:   gn,
		}
		l.Persons = append(l.Persons, p)
	}
	// извлекаем данные о подгруппах группы для групировки в структуру
	rows, err = r.store.db.Query("SELECT groupname FROM groups WHERE mothergroup = $1", gn)
	if err != nil {
		r.logger.Printf("Error to get groupname from db: %s", err)
		return nil, err
	}
	defer rows.Close()
	var s string
	var a []subgroups
	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			r.logger.Printf("trouble with rows.Next: %s", err)
			return nil, err
		}
		// извлекаем данные о людях в подгруппах для группировки в структуру
		row, err := r.store.db.Query(q, s)
		if err != nil {
			r.logger.Printf("Error to get person data from db: %s", err)
			return nil, err
		}
		defer row.Close()
		for row.Next() {
			if err := row.Scan(&id, &n, &sur, &y); err != nil {
				r.logger.Printf("trouble with rows.Next: %s", err)
				return nil, err
			}
			p = &domain.Person{
				ID:          id,
				Name:        n,
				Surname:     sur,
				YearOfBirth: y,
				GroupName:   s,
			}
			p1 := append(sub.Persons, p)
			sub := subgroups{
				Group:   s,
				Persons: p1,
			}
			a = append(a, sub)
		}
	}
	l = listPerson{
		Group:     gn,
		Persons:   l.Persons,
		Subgroups: a,
	}

	jsonData, err = json.Marshal(l)
	if err != nil {
		r.logger.Info("Error with marshaling to json", err)
		return nil, err
	}
	//fmt.Println(jsonData)
	r.logger.Trace("Try to get list of persons", q)
	return jsonData, nil
}
