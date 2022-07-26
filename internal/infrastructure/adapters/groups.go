package adapters

import (
	"net/http"

	"github.com/gorilla/mux"
)

type GroupHandler struct {
	gm GroupManager
	r  *mux.Router
}

type GroupManager interface {
	MakeGroup()
	UpdateGroup()
	DeleteGroup()
	ListAllGroups()
}

func NewGroupHandler(g GroupManager) *GroupHandler {
	return &GroupHandler{gm: g}
}

func (s *GroupHandler) RegisterGH(router *mux.Router) {
	router.HandleFunc("/creategroup", s.createGroupHandler)
}

func (gh *GroupHandler) ListAllGroups() http.Handler {
	return nil
}

func (gh *GroupHandler) createGroupHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("create group"))
}

func (gh *GroupHandler) UpdateGroupHandler() http.Handler {
	gh.gm.UpdateGroup()
	return nil
}

func (gh *GroupHandler) DeleteGroupHandler() http.Handler {
	gh.gm.DeleteGroup()
	return nil
}
