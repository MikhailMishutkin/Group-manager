package usecases

import (
	"github.com/MikhailMishutkin/Test_MediaSoft/internal/domain"
	"github.com/sirupsen/logrus"
)

//инъекция интерфейса репозитория групп
type GroupManage struct {
	repo   GroupRepository
	logger *logrus.Logger
}

// ...
type GroupRepository interface {
	CreateGroup(g *domain.Group) error
	CreateSubGroup(g *domain.Group) error
	UpdateGroup(id int, gn string, sg bool, mg string) error
	DeleteGroup(gn string) error
	GetList() (jsonData []byte, err error)
}

// конструктор
func NewGroupManage(r GroupRepository) *GroupManage {
	return &GroupManage{repo: r}
}

// создание группы или подгруппы
func (gm *GroupManage) CreateGroup(g *domain.Group) error {
	//fmt.Println(g)
	if g.Subgroup {
		err := gm.repo.CreateSubGroup(g)
		return err
	} else {
		err := gm.repo.CreateGroup(g)
		return err
	}
}

// обновление группы
func (gm *GroupManage) UpdateGroup(g *domain.Group) error {
	err := gm.repo.UpdateGroup(g.ID, g.GroupName, g.Subgroup, g.MotherGroup)
	return err
}

// удаление группы
func (gm *GroupManage) DeleteGroup(g *domain.Group) error {

	err := gm.repo.DeleteGroup(g.GroupName)
	return err

}

// выводим список групп с количеством участников этой группы и общим количеством со всеми подгруппами
func (gm *GroupManage) ListGroups() []byte {
	gm.logger = logrus.New()
	js, err := gm.repo.GetList()
	if err != nil {
		gm.logger.Info(err)
	}

	return js
}
