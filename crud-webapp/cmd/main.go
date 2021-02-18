package main

import (
	"github.com/knowit/gogo-gopher/crud-webapp/pkg/config"
	"github.com/knowit/gogo-gopher/crud-webapp/pkg/runenv"
	"log"
)

func main() {
	log.Println("Hello, world!")
	runenv.PrintAllEnvs()

	initializedConfig := config.New()
	log.Println("Config: ", initializedConfig)
}
