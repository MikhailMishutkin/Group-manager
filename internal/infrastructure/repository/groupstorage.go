package repository

import (
	"fmt"

	"github.com/MikhailMishutkin/Test_MediaSoft/internal/domain"
	"github.com/lib/pq"
)

type GroupRepository struct {
	store *Store
}

func NewGroupRepository(store *Store) *GroupRepository {
	return &GroupRepository{
		store: store,
	}
}

// ...
func (gr *GroupRepository) MakeGroup(g *domain.Group) (*domain.Group, error) {
	var gs []string
	//складываем названия групп в массив
	err := gr.store.db.QueryRow(`SELECT groupname FROM groups`).Scan(pq.Array(gs))
	if err != nil {
		fmt.Printf("Error to get groupnames from db: %s", err)
		//return nil, err
	}
	// проверяем существование группы в бд
	for _, v := range gs {
		if g.GroupName == v {
			err = fmt.Errorf("group is already exist")
			return nil, err
		} else {
			continue
		}
	}

	if err := gr.store.db.QueryRow(
		"INSERT INTO groups (groupname) VALUES ($1) RETURNING id",
		g.GroupName,
	).Scan(&g.ID); err != nil {
		return nil, err
	}

	return g, err

}

// ...
func (gr *GroupRepository) UpdateGroup() {

}

// ...
func (gr *GroupRepository) DeleteGroup() {

}

// ...
func (gr *GroupRepository) GetListAll() {

}

// ...
func (gr *GroupRepository) GetList() {

}
