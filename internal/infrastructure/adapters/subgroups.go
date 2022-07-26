package adapters

import (
	"net/http"

	"github.com/gorilla/mux"
)

type SubGroupHandler struct {
	sgm SubGroupManager
	r   *mux.Router
}

// SubGroup usecase
type SubGroupManager interface {
	MakeSubGroup()
	UpdateSubGroup()
	DeleteSubGroup()
}

func NewSubGroupHandler(sg SubGroupManager) *SubGroupHandler {
	return &SubGroupHandler{sgm: sg}
}

func (s *SubGroupHandler) RegisterSGH(router *mux.Router) {
	router.HandleFunc("/createsubgroup", s.createSubGroupHandler)
}

func (gh *SubGroupHandler) createSubGroupHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("create subgroup"))
}

func (sgh *SubGroupHandler) UpdateSubGroupHandler() http.Handler {
	sgh.sgm.UpdateSubGroup()
	return nil
}

func (sgh *SubGroupHandler) DeleteSubGroupHandler() http.Handler {
	sgh.sgm.DeleteSubGroup()
	return nil
}
