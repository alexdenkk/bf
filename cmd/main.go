package main

import (
	"os"
	"log"
	"alexdenkk/bf/internal/bf"
)

func main() {
	compiler := bf.New()

	if len(os.Args) > 1 {
		err := compiler.CompileFile(os.Args[1])

		if err != nil {
			log.Fatal(err)
		}
	}
}
