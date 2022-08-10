package usecases

import (
	"log"

	"github.com/MikhailMishutkin/Test_MediaSoft/internal/domain"
	"github.com/sirupsen/logrus"
)

//usecase struct
type GroupManage struct {
	repo   GroupRepository
	logger *logrus.Logger
}

type GroupRepository interface {
	CreateGroup(g *domain.Group) error
	CreateSubGroup(g *domain.Group) error
	UpdateGroup(id int, gn string, sg bool, mg string) error
	DeleteGroup(gn string) error
	GetListAll()
	GetList() (jsonData []byte, err error)
}

func NewGroupManage(r GroupRepository) *GroupManage {
	return &GroupManage{repo: r}
}

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

func (gm *GroupManage) ListGroups() []byte {
	js, err := gm.repo.GetList()
	if err != nil {
		log.Fatal(err)
	}

	return js
}

func (gm *GroupManage) UpdateGroup(g *domain.Group) error {
	err := gm.repo.UpdateGroup(g.ID, g.GroupName, g.Subgroup, g.MotherGroup)
	return err
}

func (gm *GroupManage) DeleteGroup(g *domain.Group) error {

	err := gm.repo.DeleteGroup(g.GroupName)
	return err

}

// func (uc *GroupManage) Make(ctx context.Context, p domain.Group) (domain.Group, error) {
// 	group, err := uc.webApi.MakeGroup(p)
// 	if err != nil {
// 		return domain.Group{}, fmt.Errorf("error to make a new Group in usecase: %w", err)
// 	}

// 	err = uc.repo.MakeGroup(context.Background(), group)
// 	if err != nil {
// 		return domain.Group{}, fmt.Errorf("error to make a new Group in repo : %w", err)
// 	}
// 	return group, nil
// }

// // вывод списка людей общий
// func (uc *GroupManage) ViewGroupsListAll(ctx context.Context) ([]domain.Group, error) {
// 	list, err := uc.repo.GetListAll(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("error to get list in usecase method: %w", err)
// 	}
// 	return list, nil
// }

// // вывод списка людей только в группе
// func (uc *GroupManage) ViewGroupsList(ctx context.Context) ([]domain.Group, error) {
// 	list, err := uc.repo.GetList(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("error to get list in usecase method: %w", err)
// 	}
// 	return list, nil
// }

// //доделать
// func (m *GroupManage) UpdateGroup() {
// 	return
// }

// func (m *GroupManage) DeleteGroup() {
// 	return
// }
