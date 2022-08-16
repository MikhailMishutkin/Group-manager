package adapters

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MikhailMishutkin/Test_MediaSoft/internal/domain"
	"github.com/gorilla/mux"
)

// структура с инъекцией из юзкейса
type GroupHandler struct {
	gm GroupManager
	r  *mux.Router
}

// ...
type GroupManager interface {
	CreateGroup(g *domain.Group) error
	UpdateGroup(g *domain.Group) error
	DeleteGroup(g *domain.Group) error
	ListGroups() []byte
}

// конструктор
func NewGroupHandler(g GroupManager) *GroupHandler {
	return &GroupHandler{gm: g}
}

// регистрация API
func (gh *GroupHandler) RegisterGH(router *mux.Router) {
	router.HandleFunc("/creategroup", gh.CreateGroupHandler).Methods("POST")
	router.HandleFunc("/listgroups", gh.ListGroupHandler).Methods("GET")
	router.HandleFunc("/updategroup", gh.UpdateGroupHandler).Methods("PUT")
	router.HandleFunc("/deletegroup", gh.DeleteGroupHandler).Methods("DELETE")
}

// создание новой группы или подгруппы
func (gh *GroupHandler) CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var g *domain.Group
	err := json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		w.WriteHeader(400)
		log.Fatal("user data error", err)
	}

	err = gh.gm.CreateGroup(g)

	if err != nil {
		JSONError(406, "not acceptable", err.Error(), w)
		return
	}

	json.NewEncoder(w).Encode(err)
}

// обновление группы
func (gh *GroupHandler) UpdateGroupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var g *domain.Group
	err := json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		w.WriteHeader(400)
		log.Fatal("user data error", err)
	}

	err = gh.gm.UpdateGroup(g)

	if err != nil {
		JSONError(406, "not acceptable", err.Error(), w)
		return
	}

	json.NewEncoder(w).Encode(err)

}

// удаление группы
func (gh *GroupHandler) DeleteGroupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var g *domain.Group
	err := json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		w.WriteHeader(400)
		log.Fatal("user data error", err)
	}

	err = gh.gm.DeleteGroup(g)

	if err != nil {
		JSONError(406, "not acceptable", err.Error(), w)
		return
	}

	json.NewEncoder(w).Encode(err)
}

// отображает список групп и количество участников в этой группе, как чисто
// в данной группе, так и общее количество вместе с дочерними группами
func (gh *GroupHandler) ListGroupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list := gh.gm.ListGroups()
	w.Write(list)
}
