package usecases

//в юзкейсе описывается не каким образом программа делает что-либо,
//а что именно она делает
import (
	"context"
	"log"

	"github.com/MikhailMishutkin/Test_MediaSoft/internal/domain"
	"github.com/sirupsen/logrus"
)

// type PersonManager interface {
//
// }

type PersonManage struct {
	repo   PersonRepository
	logger *logrus.Logger
}

type PersonRepository interface {
	CreatePerson(p *domain.Person) (*domain.Person, error)
	UpdatePerson(id int, gn string) error
	DeletePerson(id int) error
	GetListAll()
	GetList() ([]byte, error)
}

func NewPersonManage(r PersonRepository) *PersonManage {
	return &PersonManage{repo: r}
}

//создание профиля человека
func (pm *PersonManage) CreatePerson(c *domain.Person) error {
	//fmt.Println(c)
	_, err := pm.repo.CreatePerson(c)
	return err
}

// вывод списка людей общий
func (pm *PersonManage) ViewPersonsListAll() []byte {
	pm
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

//...
func (pm *PersonManage) UpdatePerson(p *domain.Person) error {

	err := pm.repo.UpdatePerson(p.ID, p.GroupName)
	return err
}

// ...
func (pm *PersonManage) DeletePerson(p *domain.Person) error {
	err := pm.repo.DeletePerson(p.ID)
	return err
}
