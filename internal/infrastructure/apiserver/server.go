package apiserver

import (
	"net/http"

	"github.com/MikhailMishutkin/Test_MediaSoft/internal/infrastructure/adapters"
	"github.com/MikhailMishutkin/Test_MediaSoft/internal/infrastructure/repository"
	"github.com/MikhailMishutkin/Test_MediaSoft/internal/usecases"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	logger *logrus.Logger
	router *mux.Router
	store  repository.Storer

	ph  *adapters.PersonHandler
	gh  *adapters.GroupHandler
	sgh *adapters.SubGroupHandler
}

func newServer(store repository.Storer) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,

		ph:  &adapters.PersonHandler{},
		gh:  &adapters.GroupHandler{},
		sgh: &adapters.SubGroupHandler{},
	}

	repository := store.Person()
	useCase := usecases.NewPersonManage(repository)
	c := adapters.NewUserHandler(useCase)
	c.RegisterPH(s.router)
	s.logger.Info("starting api server")
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
