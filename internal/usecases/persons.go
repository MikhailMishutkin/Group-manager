package usecases

//в юзкейсе описывается не каким образом программа делает что-либо,
//а что именно она делает
import (
	"github.com/MikhailMishutkin/Test_MediaSoft/internal/domain"
	"github.com/sirupsen/logrus"
)

// инъекция интерфейса репозитория
type PersonManage struct {
	repo   PersonRepository
	logger *logrus.Logger
}

// интерфейс
type PersonRepository interface {
	CreatePerson(p *domain.Person) (*domain.Person, error)
	UpdatePerson(id int, gn string) error
	DeletePerson(id int) error
	GetList(gn string) ([]byte, error)
}

// конструктор
func NewPersonManage(r PersonRepository) *PersonManage {
	return &PersonManage{repo: r}
}

// метод создания профиля человека
func (pm *PersonManage) CreatePerson(c *domain.Person) error {
	_, err := pm.repo.CreatePerson(c)
	return err
}

// метод обновления профиля человека по id, обновляем только группу
func (pm *PersonManage) UpdatePerson(p *domain.Person) error {
	err := pm.repo.UpdatePerson(p.ID, p.GroupName)
	return err
}

// удаляем по id
func (pm *PersonManage) DeletePerson(p *domain.Person) error {
	err := pm.repo.DeletePerson(p.ID)
	return err
}

// вывод списка людей в группе и подгруппах
func (pm *PersonManage) ViewPersonsListAll(p *domain.Group) []byte {
	pm.logger = logrus.New()
	n := p.GroupName
	js, err := pm.repo.GetList(n)
	if err != nil {
		pm.logger.Info(err)
	}

	return js
}
