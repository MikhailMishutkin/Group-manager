package config

import "os"

type Config struct {
	BindAddr string
	Host     string
	Port     string
	User     string
	Password string
	NameDB   string
}

func NewConfig() *Config {
	return &Config{
		BindAddr: os.Getenv("BindAddr"),
		Host:     os.Getenv("Host"),
		Port:     os.Getenv("Port"),
		User:     os.Getenv("User"),
		Password: os.Getenv("Password"),
		NameDB:   os.Getenv("NameDB"),
	}
}
