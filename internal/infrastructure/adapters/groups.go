package adapters

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MikhailMishutkin/Test_MediaSoft/internal/domain"
	"github.com/gorilla/mux"
)

// ...
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

// ...
func NewGroupHandler(g GroupManager) *GroupHandler {
	return &GroupHandler{gm: g}
}

// ...
func (gh *GroupHandler) RegisterGH(router *mux.Router) {
	router.HandleFunc("/creategroup", gh.CreateGroupHandler).Methods("POST")
	router.HandleFunc("/listgroups", gh.ListGroupHandler).Methods("GET")
	router.HandleFunc("/listgrougswithsub", gh.ListGroupWSHandler).Methods("GET")
	router.HandleFunc("/updategroup", gh.UpdateGroupHandler).Methods("PUT")
	router.HandleFunc("/deletegroup", gh.DeleteGroupHandler).Methods("DELETE")
}

// ...
func (gh *GroupHandler) CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var g *domain.Group
	//_ = json.NewDecoder(r.Body).Decode(&p)
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

// ...
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

// ...
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

// ...
func (gh *GroupHandler) ListGroupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list := gh.gm.ListGroups()
	w.Write(list)
}

// ?????????
func (gh *GroupHandler) ListGroupWSHandler(w http.ResponseWriter, r *http.Request) {

}
