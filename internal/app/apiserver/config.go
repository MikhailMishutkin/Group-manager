package apiserver

import "github.com/MikhailMishutkin/Test_MediaSoft/internal/infrastructure/repository"

type Config struct {
	BindAddr string          `toml:"bind_addr"`
	LogLevel string          `toml:"log_level"`
	DBUrl    repository.Data `toml:"data"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
