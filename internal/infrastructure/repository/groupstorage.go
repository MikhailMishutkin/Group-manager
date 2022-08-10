package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/MikhailMishutkin/Test_MediaSoft/internal/domain"
	"github.com/sirupsen/logrus"
)

type GroupRepository struct {
	store  *Store
	logger *logrus.Logger
}

func NewGroupRepository(store *Store) *GroupRepository {
	return &GroupRepository{
		store: store,
	}
}

// ...
func (gr *GroupRepository) CreateGroup(g *domain.Group) error {
	gr.logger = logrus.New()

	//извлекаем значения существующих групп
	rows, err := gr.store.db.Query("SELECT groupname FROM groups")
	if err != nil {
		gr.logger.Printf("Error to get groupnames from db: %s", err)
		return err
	}
	defer rows.Close()

	_, err = checkGroupNamesExst(rows, g)
	if err != nil {
		return err
	}

	// создаем группу верхнего уровня
	if err := gr.store.db.QueryRow(
		"INSERT INTO groups (groupname, members) VALUES ($1, $2) RETURNING id",
		g.GroupName,
		g.Members,
	).Scan(&g.ID); err != nil {
		return err
	}

	return err
}

// ...
func (gr *GroupRepository) CreateSubGroup(g *domain.Group) error {
	gr.logger = logrus.New()
	var gs []string //массив с названиями групп

	//извлекаем значения существующих групп для проверки существования материнской группы
	rows, err := gr.store.db.Query("SELECT groupname FROM groups")
	if err != nil {
		gr.logger.Printf("Error to get groupnames from db: %s", err)
		return err
	}
	defer rows.Close()
	gs, err = checkGroupNamesExst(rows, g)
	if err != nil {
		return err
	}

	//извлекаем значения подгрупп в группе
	rows, err = gr.store.db.Query("SELECT groupname FROM groups WHERE mothergroup = $1", g.MotherGroup)
	if err != nil {
		gr.logger.Printf("Error to get groupnames from db: %s", err)
		return err
	}
	defer rows.Close()

	_, err = checkGroupNamesExst(rows, g)
	if err != nil {
		return err
	}

	// проверяем есть ли материнская группа
	for _, v := range gs {
		if g.MotherGroup == v {
			// создаем подгруппу
			if err := gr.store.db.QueryRow(
				"INSERT INTO groups (groupname, members, mothergroup) VALUES ($1, $2, $3) RETURNING id",
				g.GroupName,
				g.Members,
				g.MotherGroup,
			).Scan(&g.ID); err != nil {
				return err
			}

			return err
		} else {
			continue
		}
	}

	err = fmt.Errorf("no such mothergroup exist")
	gr.logger.Info(err)
	return err
}

// TODO
func (gr *GroupRepository) GetList() (jsonData []byte, err error) {
	gr.logger = logrus.New()
	q := `
	SELECT json_agg(s) FROM (
		SELECT groupname
		from groups 
	) s;`
	if err := gr.store.db.QueryRow(q).Scan(&jsonData); err != nil {
		gr.logger.Printf("Error GetList of persons: %s", err)
		return nil, err
	}
	gr.logger.Tracef("Try to get list of persons", q)
	return jsonData, nil
}

// обновление группы по id: можно поменять имя, если нет участников; можно поменять уровень вверх и вниз
func (gr *GroupRepository) UpdateGroup(id int, gn string, sg bool, mg string) error {
	gr.logger = logrus.New()
	var name string
	var q int
	err := gr.store.db.QueryRow("SELECT groupname FROM groups WHERE id = $1", id).Scan(&name)
	if err != nil {
		gr.logger.Printf("Error to get groupname from db: %s", err)
		return err
	}

	err = gr.store.db.QueryRow("SELECT members FROM groups WHERE id = $1", id).Scan(&q)
	if err != nil {
		gr.logger.Printf("Error to get groupname from db: %s", err)
		return err
	}

	switch {
	case gn != name && q > 0:
		return errors.New("cannot update name, group has a members")
	case gn != name:
		_, err := gr.store.db.Exec(`UPDATE groups SET groupname = $2 WHERE id = $1`, id, gn)
		if err != nil {
			gr.logger.Infof("error to update group: %s", err)
		}
		return nil
	case gn == name && sg:
		_, err := gr.store.db.Exec(`UPDATE groups SET mothergroup = $2 WHERE id = $1`, id, mg)
		if err != nil {
			gr.logger.Infof("error to update group: %s", err)
		}
		return nil
	case gn == name && !sg:
		_, err := gr.store.db.Exec(`UPDATE groups SET mothergroup = $2 WHERE id = $1`, id, mg)
		if err != nil {
			gr.logger.Infof("error to update group: %s", err)
		}
		return nil
	}

	return err
}

// ...
func (gr *GroupRepository) DeleteGroup(gn string) error {
	gr.logger = logrus.New()
	var q int
	err := gr.store.db.QueryRow("SELECT members FROM groups WHERE groupname = $1", gn).Scan(&q)
	if err != nil {
		gr.logger.Printf("Error to get groupname from db: %s", err)
		return err
	}

	if q > 0 {
		return errors.New("cannot delete, group has a members")
	} else {
		_, err = gr.store.db.Exec(`DELETE FROM groups where groupname = $1`, gn)
		if err != nil {
			gr.logger.Info("error to delete person: %s", err)
			return err
		}
	}

	return nil
}

// ...
func (gr *GroupRepository) GetListAll() {

}

// ...
func checkGroupNamesExst(rows *sql.Rows, g *domain.Group) ([]string, error) {
	logger := logrus.New()
	var gs []string //массив с названиями групп
	var s string

	//складываем имена групп в массив
	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			logger.Printf("trouble with rows.Next: %s", err)
			return nil, err
		}
		gs = append(gs, s)
	}

	// проверяем существование группы в бд
	for _, v := range gs {
		if g.GroupName == v {
			err := fmt.Errorf("group is already exist")
			return nil, err
		} else {
			continue
		}
	}

	return gs, nil
}

// //извлекаем значения существующих групп
// if g.Subgroup {
// 	rows, err := gr.store.db.Query("SELECT groupname FROM groups WHERE COLUMN mothergroup = $1", g.MotherGroup)
// 	if err != nil {
// 		gr.logger.Printf("Error to get groupnames from db: %s", err)
// 		return nil
// 	}
// 	defer rows.Close()

// 	//складываем группы в массив
// 	for rows.Next() {
// 		if err := rows.Scan(&s); err != nil {
// 			gr.logger.Printf("trouble with rows.Next: %s", err)
// 			return nil
// 		}
// 		gs = append(gs, s)
// 	}
// 	return gs
// } else {
// 	rows, err := gr.store.db.Query("SELECT groupname FROM groups")
// 	if err != nil {
// 		gr.logger.Printf("Error to get groupnames from db: %s", err)
// 		return nil
// 	}
// 	defer rows.Close()
// 	//складываем группы в массив
// 	for rows.Next() {
// 		if err := rows.Scan(&s); err != nil {
// 			gr.logger.Printf("trouble with rows.Next: %s", err)
// 			return nil
// 		}
// 		gs = append(gs, s)
// 	}
// 	return gs
// }
// //складываем группы в массив
// for rows.Next() {
// 	if err := rows.Scan(&s); err != nil {
// 		return g, err
// 	}
// 	gs = append(gs, s)
// }
// // проверяем существование группы в бд
// for _, v := range gs {
// 	if g.GroupName == v {
// 		err = fmt.Errorf("group is already exist")
// 		return nil, err
// 	} else {
// 		continue
// 	}
// }

// for _, v := range gs {
// 	if g.MotherGroup == v {
// 		for _, v := range gs {
// 			if g.GroupName == v {
// 				err := fmt.Errorf("group is already exist")
// 				gr.logger.Error("group is already exist")
// 				return err
// 			} else {
// 				continue
// 			}
// 		}
// 		return nil
// 	} else {
// 		continue
// 	}
// }
// err := fmt.Errorf("no such mothergroup exist")
// gr.logger.Error("no such mothergroup exist")
// return err
