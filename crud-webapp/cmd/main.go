package main

import (
	"log"

	"github.com/knowit/gogo-gopher/crud-webapp/config"
	"github.com/knowit/gogo-gopher/crud-webapp/runenv"
)

func main() {
	log.Println("Hello, world!")
	runenv.PrintAllEnvs()

	initializedConfig := config.New()
	log.Println("Config: ", initializedConfig)
}
