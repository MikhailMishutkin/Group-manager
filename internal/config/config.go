package config

type Config struct {
	BindAddr string   `toml:"bind_addr"`
	LogLevel string   `toml:"log_level"`
	DB       Database `toml:"database"`
}

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	NameDB   string
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
