package apiserver

import (
	"database/sql"
	"fmt"
	"net/http"

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
