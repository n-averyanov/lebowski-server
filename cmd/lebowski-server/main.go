package main

import (
	"github.com/BurntSushi/toml"
	"flag"
	"log"

	apiserver "github.com/n-averyanov/lebowski-server/internal/app/lebowski-server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if(err != nil) {
		log.Fatal(err)
	}

	server := apiserver.New(config)
	if error := server.Start(); error != nil {
		log.Fatal(error)
	}
}