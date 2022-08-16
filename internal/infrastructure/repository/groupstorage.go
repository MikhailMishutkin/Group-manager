package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/MikhailMishutkin/Test_MediaSoft/internal/domain"
	"github.com/sirupsen/logrus"
)

// структура для вывода списка
type listGroups struct {
	Id        int    `json:"id"`
	GroupName string `json:"group_name"`
	MembInG   int    `json:"members_in_group"`
	MembAll   int    `json:"members_in_total_with_subgroups"`
}

// структура с инъекцией настроек хранилища
type GroupRepository struct {
	store  *Store
	logger *logrus.Logger
}

// конструктор
func NewGroupRepository(store *Store) *GroupRepository {
	return &GroupRepository{
		store: store,
	}
}

// создаём группу
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
		"INSERT INTO groups (groupname, members, subgroup) VALUES ($1, $2, $3) RETURNING id",
		g.GroupName,
		g.Members,
		g.Subgroup,
	).Scan(&g.ID); err != nil {
		return err
	}

	return err
}

// создаём подгруппу
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
				"INSERT INTO groups (groupname, members, subgroup, mothergroup) VALUES ($1, $2, $3, $4) RETURNING id",
				g.GroupName,
				g.Members,
				g.Subgroup,
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

	// TODO: прописать вариант отказа изменения уровня, если у группы есть подгруппы
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
		_, err := gr.store.db.Exec(`UPDATE groups SET mothergroup = $2, subgroup = true WHERE id = $1`, id, mg)
		if err != nil {
			gr.logger.Infof("error to update group: %s", err)
		}
		return nil
	case gn == name && !sg:
		_, err := gr.store.db.Exec(`UPDATE groups SET mothergroup = $2, subgroup = false WHERE id = $1`, id, mg)
		if err != nil {
			gr.logger.Infof("error to update group: %s", err)
		}
		return nil
	}

	return err
}

// удаление группы, нельзя удалить, если в группе есть участники или подгруппы
func (gr *GroupRepository) DeleteGroup(gn string) error {
	gr.logger = logrus.New()
	var q int
	var b bool
	err := gr.store.db.QueryRow("SELECT members, subgroup FROM groups WHERE groupname = $1", gn).Scan(&q, &b)
	if err != nil {
		gr.logger.Printf("Error to get groupname from db: %s", err)
		return err
	}
	// проверка на участников
	if q > 0 {
		return errors.New("cannot delete, group has a members")
	} else {
		_, err = gr.store.db.Exec(`DELETE FROM groups where groupname = $1`, gn)
		if err != nil {
			gr.logger.Info("error to delete group: %s", err)
			return err
		}
	}
	// проверка на подгруппы
	var s string
	if !b {
		err := gr.store.db.QueryRow("SELECT groupname FROM groups WHERE mothergroup = $1", gn).Scan(&s)
		if err != nil {
			gr.logger.Printf("Error to get groupname from db: %s", err)
			return err
		}
		if len(s) > 0 {
			return errors.New("cannot delete, group has a subgroup")
		} else {
			_, err = gr.store.db.Exec(`DELETE FROM groups where groupname = $1`, gn)
			if err != nil {
				gr.logger.Info("error to delete group: %s", err)
				return err
			}

		}
	}

	return nil
}

// отображает список групп и количество участников в этой группе, как чисто в данной группе,
//так и общее количество вместе с дочерними группами
func (gr *GroupRepository) GetList() (jsonData []byte, err error) {
	var s string
	var l listGroups
	gr.logger = logrus.New()

	// делаем выборку материнских групп из списка
	rows, err := gr.store.db.Query("SELECT groupname FROM groups WHERE subgroup = false")
	if err != nil {
		gr.logger.Printf("Error to get groupnames from db: %s", err)
		return nil, err
	}
	defer rows.Close()
	// сканируем группу в s
	var listAll []listGroups
	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			gr.logger.Printf("trouble with rows.Next: %s", err)
			return nil, err
		}
		//получаем id и количество участников группы
		var id, m int
		err = gr.store.db.QueryRow("SELECT id, members FROM groups WHERE groupname = $1", s).Scan(&id, &m)
		if err != nil {
			gr.logger.Printf("Error to get id and members from db: %s", err)
			return nil, err
		}
		// делаем выборку количества участников из подгрупп
		members, err := gr.store.db.Query("SELECT members FROM groups WHERE mothergroup = $1", s)
		if err != nil {
			gr.logger.Printf("Error to get groupnames from db: %s", err)
			return nil, err
		}
		defer members.Close()
		// суммируем участников подгрупп
		var q, sumMemb int
		for members.Next() {
			if err := members.Scan(&q); err != nil {
				gr.logger.Printf("trouble with rows.Next: %s", err)
				return nil, err
			}
			sumMemb = sumMemb + q
		}
		sumMemb = sumMemb + m
		// складываем полученные данные в структуру и сериализируем в json
		l = listGroups{
			Id:        id,
			GroupName: s,
			MembInG:   m,
			MembAll:   sumMemb,
		}
		listAll = append(listAll, l)
	}
	jsonData, err = json.Marshal(listAll)
	if err != nil {
		gr.logger.Info("Error with marshaling to json", err)
		return nil, err
	}
	//fmt.Println(string(jsonData))
	return jsonData, nil
}

// вспомогательная функция проверки существования группы в бд
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
