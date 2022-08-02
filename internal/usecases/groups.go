package usecases

import (
	"fmt"

	"github.com/MikhailMishutkin/Test_MediaSoft/internal/domain"
)

//import "github.com/MikhailMishutkin/Test_MediaSoft/internal/domain"

//usecase struct
type GroupManage struct {
	repo GroupRepository
	//webApi GroupWebApi
}

type GroupRepository interface {
	MakeGroup(g *domain.Group) (*domain.Group, error)
	UpdateGroup()
	DeleteGroup()
	GetListAll()
	GetList()
}

func NewGroupManage(r GroupRepository) *GroupManage {
	return &GroupManage{repo: r}
}

func (gm *GroupManage) CreateGroup(g *domain.Group) {
	fmt.Println(g)
	gm.repo.MakeGroup(g)
}

func (gm *GroupManage) ListGroup() {

}

func (gm *GroupManage) UpdateGroup() {
	gm.repo.DeleteGroup()
}

func (gm *GroupManage) DeleteGroup() {
	gm.repo.DeleteGroup()

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
