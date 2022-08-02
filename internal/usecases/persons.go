package usecases

//в юзкейсе описывается не каким образом программа делает что-либо,
//а что именно она делает
import (
	"context"
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
	UpdatePerson(id int, gn string)
	DeletePerson(id int)
	GetListAll()
	GetList() ([]byte, error)
}

func NewPersonManage(r PersonRepository) *PersonManage {
	return &PersonManage{repo: r}
}

//создание профиля человека
func (pm *PersonManage) CreatePerson(c *domain.Person) error {
	//fmt.Println(c)
	_, err := pm.repo.MakePerson(c)
	return err
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
func (pm *PersonManage) UpdatePerson(p *domain.Person) {

	pm.repo.UpdatePerson(p.ID, p.GroupName)
}

func (pm *PersonManage) DeletePerson(p *domain.Person) {
	pm.repo.DeletePerson(p.ID)
	return
}
