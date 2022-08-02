package adapters

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MikhailMishutkin/Test_MediaSoft/internal/domain"
	"github.com/gorilla/mux"
)

type GroupHandler struct {
	gm GroupManager
	r  *mux.Router
}

type GroupManager interface {
	CreateGroup(g *domain.Group)
	UpdateGroup()
	DeleteGroup()
	ListGroup()
}

func NewGroupHandler(g GroupManager) *GroupHandler {
	return &GroupHandler{gm: g}
}

func (gh *GroupHandler) RegisterGH(router *mux.Router) {
	router.HandleFunc("/creategroup", gh.CreateGroupHandler).Methods("POST")
	router.HandleFunc("/listgroups", gh.ListGroupHandler).Methods("GET")
	router.HandleFunc("/listgrougswithsub", gh.ListGroupWSHandler).Methods("GET")
	router.HandleFunc("/updategroup", gh.UpdateGroupHandler).Methods("PUT")
	router.HandleFunc("/deletegroup", gh.DeleteGroupHandler).Methods("DELETE")
}

func (gh *GroupHandler) ListGroupHandler(w http.ResponseWriter, r *http.Request) {
}

func (gh *GroupHandler) CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var g *domain.Group
	//_ = json.NewDecoder(r.Body).Decode(&p)
	err := json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		w.WriteHeader(400)
		log.Fatal("user data error", err)
	}
	//fmt.Println(g)
	gh.gm.CreateGroup(g)

	json.NewEncoder(w).Encode(g)
}

func (gh *GroupHandler) ListGroupWSHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("create group"))
}

func (gh *GroupHandler) UpdateGroupHandler(w http.ResponseWriter, r *http.Request) {
	gh.gm.UpdateGroup()
}

func (gh *GroupHandler) DeleteGroupHandler(w http.ResponseWriter, r *http.Request) {
	gh.gm.DeleteGroup()
}
