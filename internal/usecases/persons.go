package usecases

//в юзкейсе описывается не каким образом программа делает что-либо,
//а что именно она делает
import (
	"context"
	"fmt"
	"log"

	"github.com/MikhailMishutkin/Test_MediaSoft/internal/domain"
)

// type PersonManager interface {
//
// }

type PersonManage struct {
	repo PersonRepository
}

type PersonRepository interface {
	MakePerson(p *domain.Person) (*domain.Person, error)
	UpdatePerson()
	DeletePerson()
	GetListAll()
	GetList() ([]byte, error)
}

func NewPersonManage(r PersonRepository) *PersonManage {
	return &PersonManage{repo: r}
}

//создание профиля человека
func (pm *PersonManage) CreatePerson(c *domain.Person) {
	fmt.Println(c)
	pm.repo.MakePerson(c)
	// person, err := pm.repo.MakePerson()
	// if err != nil {
	// 	return domain.Person{}, fmt.Errorf("error to make a new person in usecase: %w", err)
	// }

	// err = uc.repo.MakePerson(context.Background(), person)
	// if err != nil {
	// 	return domain.Person{}, fmt.Errorf("error to make a new person in repo : %w", err)
	// }
	// return person, nil
}

// вывод списка людей общий
func (pm *PersonManage) ViewPersonsListAll() []byte {
	js, err := pm.repo.GetList()
	if err != nil {
		log.Fatal(err)
	}

	return js
}

// 	list, err := uc.repo.GetListAll(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("error to get list in usecase method: %w", err)
// 	}
// 	return list, nil
// }

// вывод списка людей только в группе
func (pm *PersonManage) ViewPersonsList(ctx context.Context) {
	// list, err := pm.repo.GetListAll()
	// if err != nil {
	// 	return nil, fmt.Errorf("error to get list in usecase method: %w", err)
	// }
	// return list, nil
}

//доделать
func (m *PersonManage) UpdatePerson() {
	return
}

func (m *PersonManage) DeletePerson() {
	return
}
