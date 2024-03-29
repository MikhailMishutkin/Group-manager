package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/MikhailMishutkin/Test_MediaSoft/internal/app/apiserver"
	"github.com/MikhailMishutkin/Test_MediaSoft/internal/config"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := config.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}

}
