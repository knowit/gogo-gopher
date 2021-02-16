package runenv

import (
	"fmt"
	"os"
)

func PrintAllEnvs() {
	fmt.Println("Printing all environment variables...")
	for _, entry := range os.Environ() {
		fmt.Println(entry)
	}
}
