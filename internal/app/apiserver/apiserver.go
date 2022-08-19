package apiserver

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/MikhailMishutkin/Test_MediaSoft/internal/config"
	"github.com/MikhailMishutkin/Test_MediaSoft/internal/infrastructure/repository"

	_ "github.com/lib/pq" // ...
	"github.com/sirupsen/logrus"
)

func Start(config *config.Config) error {

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
	var conn *config.Config
	if _, err := toml.DecodeFile("configs/apiserver.toml", &conn); err != nil {
		return nil, err
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conn.DB.Host, conn.DB.Port, conn.DB.User, conn.DB.Password, conn.DB.NameDB)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func configureLogger(config *config.Config) error {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}

	logrus.SetLevel(level)

	return nil
}
