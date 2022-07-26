package app

import (
	"github.com/MikhailMishutkin/Test_MediaSoft/internal/infrastructure/adapters"
	"github.com/MikhailMishutkin/Test_MediaSoft/internal/infrastructure/repository"
	"github.com/MikhailMishutkin/Test_MediaSoft/internal/usecases"
	"github.com/gorilla/mux"
)

func Run() {
	repository := repository.NewPersonRepository()
	useCase := usecases.NewPersonManage(repository)
	c := adapters.NewUserHandler(useCase)
	c.RegisterPH(mux.NewRouter())

}
