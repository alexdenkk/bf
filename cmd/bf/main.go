package main

import (
	"alexdenkk/bf/internal/bf"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const HELP_MSG = `
brainfuck compiler

Usage:
build <filename> - compile file

Commands:
, | input cell
. | print cell
> | next cell
< | previous cell
+ | plus 1 to cell value
- | minus to cell value
[ | open cycle
] | close cycle

*Cycles must be closed

Basic expressions:
,>,[<+>-]<.                                | summation
>,[<->-]<.                                 | subtraction
,>,[<[->>+>+<<<]>>>[-<<<+>>>]<<<>-]<[-]>>. | multiplication`

func main() {
	compiler := bf.NewCompiler()

	if len(os.Args) > 2 {
		bfCode, err := readFile(os.Args[2])

		if err != nil {
			log.Fatal(err)
		}

		goCode, err := compiler.Compile(bfCode)

		if err != nil {
			log.Fatal(err)
		}

		err = writeFile(os.Args[2][:len(os.Args[2])-3]+".go", goCode)

		if err != nil {
			log.Fatal(err)
		}

		err = compileGoFile(os.Args[2][:len(os.Args[2])-3])

		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println(HELP_MSG)
	}
}

func compileGoFile(filename string) error {
	cmd := exec.Command("go", "build", filename+".go")

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func readFile(filename string) (string, error) {
	b, err := os.ReadFile(filename)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

func writeFile(filename string, text string) error {
	return os.WriteFile(filename, []byte(text), 0644)
}
