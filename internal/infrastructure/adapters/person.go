package adapters

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MikhailMishutkin/Test_MediaSoft/internal/domain"

	//"github.com/MikhailMishutkin/Test_MediaSoft/internal/usecases"
	"github.com/gorilla/mux"
)

type PersonHandler struct {
	service PersonManager
}

// методы в юзкейсе
type PersonManager interface {
	CreatePerson(c *domain.Person) error
	ViewPersonsListAll() []byte
	UpdatePerson(p *domain.Person) error
	DeletePerson(p *domain.Person) error
}

//...
func NewUserHandler(pm PersonManager) *PersonHandler {
	return &PersonHandler{service: pm}
}

//...
func (s *PersonHandler) RegisterPH(router *mux.Router) {
	router.HandleFunc("/createperson", s.CreatePersonHandler).Methods("POST")
	// п. 4 здесь или в группах??????????
	router.HandleFunc("/listperson", s.ListPersonHandler).Methods("GET")
	router.HandleFunc("/listpersonws", s.ListPersonWithSubHandler).Methods("GET")
	router.HandleFunc("/updateperson", s.UpdatePersonHandler).Methods("PUT")
	router.HandleFunc("/deleteperson", s.DeletePersonHandler).Methods("DELETE")
}

// создаёт запись о человеке
func (u *PersonHandler) CreatePersonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var p *domain.Person
	//_ = json.NewDecoder(r.Body).Decode(&p)
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

/// отображает список людей
func (u *PersonHandler) ListPersonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fromSt := u.service.ViewPersonsListAll()
	w.Write(fromSt)

}

//???????
func (u *PersonHandler) ListPersonWithSubHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

//...
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

//...
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

// func (r PersonHandler) AddRouterPerson() {
// 	r.r.Get("/api/createPerson/").Methods("GET").Handler(r.CreatePerson())
// {
// 	r.Get("/{UserID}/all", GetArchiveIDs)
// 	r.Get("/{UserID}/{ExchangeID}", GetArchiveDetails)
// }
//r.r.HandleFunc("/api/createPerson", r.CreatePerson()).Methods("POST")
// 	s := r.r.Host("8080").Subrouter()
// 	s.HandleFunc("/person/", r.CreatePerson())
// 	//s.HandleFunc("/products/{key}", ProductHandler)
// 	//s.HandleFunc("/articles/{category}/{id:[0-9]+}"), ArticleHandler)

// 	r.r.Handle("api/listPerson/{groupId}", r.ListPerson()).Methods("GET")
// 	r.r.Handle("api/listPerson/{groupId}/subgroups", r.ListPersonWithSub()).Methods("GET")
// 	r.r.Handle("api/updatePerson/{personId}", r.UpdatePerson()).Methods("PUT")
// 	r.r.Handle("api/deletePerson/{personId}", r.DeletePerson()).Methods("DELETE")
