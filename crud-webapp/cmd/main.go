package main

import (
	"fmt"
	"github.com/knowit/gogo-gopher/crud-webapp/pkg/runenv"
)

func main() {
	fmt.Println("Hello, world!")
	runenv.PrintAllEnvs()
}
