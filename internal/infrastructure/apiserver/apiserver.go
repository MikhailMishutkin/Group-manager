package apiserver

import (
	"database/sql"
	"fmt"
	"net/http"

	//"github.com/MikhailMishutkin/Test_MediaSoft/internal/infrastructure/adapters"

	"github.com/MikhailMishutkin/Test_MediaSoft/internal/infrastructure/repository"
	_ "github.com/lib/pq" // ...
	"github.com/sirupsen/logrus"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1Vfcmrf1"
	dbname   = "testMS"
)

func Start(config *Config) error {

	db, err := newDB()
	if err != nil {
		return err
	}
	defer db.Close()
	store := repository.New(db)
	srv := newServer(store)

	// if err := s.configureStore(); err != nil {
	// 	return err
	// }

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func configureLogger(config *Config) error {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}

	logrus.SetLevel(level)

	return nil
}

// func (s *APIServer) configureStore() error {
// 	st := repository.NewPersonRepository(s.config.Repository)
// 	if err := st.Open(); err != nil {
// 		return err
// 	}

// 	s.PersonRepo = st

// 	return nil
// }

// // создаём структуру для сервера
// type APIServer struct {
// 	config *Config // поле конфиг со ссылкой на тип конфиг с прописанными параметрами

// }

// // делаем конструктор-копию
// func New(config *Config) *APIServer {
// 	return &APIServer{
// 		config: config,
// 	}
// }

// прописываем методы для структуры
// func (s *APIServer) configureRouter() {
// 	s.ph.RegisterPH(s.router)
// 	s.gh.RegisterGH(s.router)
// 	s.sgh.RegisterSGH(s.router)
// }

//
// 	sl := s.router.Host("localhost:8080").Subrouter()
// 	sl.HandleFunc("/person/", s.handleRoutes()).Methods("GET")

// 	//s.router.Get("/api/createPerson").Methods("GET").Subrouter().NewRoute().HandlerFunc(adapters.PersonManager.CreatePerson())
// 	//adapters.SubGroupHandle.AddRouterSubGroup(adapters.SubGroupHandle{})
// 	//adapters.PersonHandler.AddRouterPerson(adapters.PersonHandler{})
// 	//adapters.GroupHandler.AddRouter(adapters.GroupHandler{})
// }

// func (s *APIServer) handleTest() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		io.WriteString(w, "test")
// 	}
// }

// func (s *APIServer) handleRoutes() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w = adapters.PersonHandler.CreatePerson()
// 		io.WriteString(w, "test2")
// 	}
// }
