package apiserver

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/MikhailMishutkin/Test_MediaSoft/internal/config"
	"github.com/MikhailMishutkin/Test_MediaSoft/internal/infrastructure/repository"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq" // ...
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c := config.NewConfig()

	fmt.Println(c.Host)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.NameDB)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
