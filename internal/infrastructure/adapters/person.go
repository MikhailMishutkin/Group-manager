package adapters

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MikhailMishutkin/Test_MediaSoft/internal/domain"

	"github.com/gorilla/mux"
)

//структура с инъекцией юзкейса
type PersonHandler struct {
	service PersonManager
}

// методы в юзкейсе
type PersonManager interface {
	CreatePerson(c *domain.Person) error
	ViewPersonsListAll(p *domain.Group) []byte
	UpdatePerson(p *domain.Person) error
	DeletePerson(p *domain.Person) error
}

//конструктор
func NewUserHandler(pm PersonManager) *PersonHandler {
	return &PersonHandler{service: pm}
}

//регистрация API
func (s *PersonHandler) RegisterPH(router *mux.Router) {
	router.HandleFunc("/createperson", s.CreatePersonHandler).Methods("POST")
	router.HandleFunc("/listperson", s.ListPersonHandler).Methods("GET")
	router.HandleFunc("/updateperson", s.UpdatePersonHandler).Methods("PUT")
	router.HandleFunc("/deleteperson", s.DeletePersonHandler).Methods("DELETE")
}

// создаёт запись о человеке
func (u *PersonHandler) CreatePersonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var p *domain.Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(400)
		log.Fatal("user data error", err)
	}

	err = u.service.CreatePerson(p)

	if err != nil {
		JSONError(406, "not acceptable", err.Error(), w)
		return
	}
	json.NewEncoder(w).Encode(err)
}

// обновляет запись о человеке
func (u *PersonHandler) UpdatePersonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var p *domain.Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(400)
		log.Fatal("user data error", err)
	}

	err = u.service.UpdatePerson(p)
	if err != nil {
		JSONError(406, "not acceptable", err.Error(), w)
		return
	}
	json.NewEncoder(w).Encode(err)
}

// удаляет запись о человеке
func (u *PersonHandler) DeletePersonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var p *domain.Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(400)
		log.Fatal("user data error", err)
	}
	err = u.service.DeletePerson(p)
	if err != nil {
		JSONError(406, "not acceptable", err.Error(), w)
		return
	}
	json.NewEncoder(w).Encode(err)
	w.WriteHeader(200)
}

/// отображает список людей в определённой группе только привязанных к данной
// и список людей со всеми дочерними группами
func (u *PersonHandler) ListPersonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var p *domain.Group
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(400)
		log.Fatal("user data error", err)
	}
	list := u.service.ViewPersonsListAll(p)
	w.Write(list)

}
