package bf

import (
	"os/exec"
	"io/ioutil"
	"strings"
)

type Compiler struct {
	Tokens map[string]string
}

func New() *Compiler {
	tokens := BaseMap

	return &Compiler{
		Tokens: tokens,
	}
}

func (bf *Compiler) SetTokens(tokens TokenMap) {
	tokens.Parse()
	bf.Tokens = tokens
}

func (bf *Compiler) Parse(commands string) []string {
	parsed := []string{}

	for _, command := range commands {
		if bfCommand, ok := bf.Tokens[string(command)]; ok {
			parsed = append(parsed, bfCommand)
		}
	}

	println("parsed")

	return parsed
}

func (bf *Compiler) Translate(commands []string) (string, error) {
	end := BEGIN
	tabs := 1

	for _, command := range commands {

		end = end + "\n" + strings.Repeat(TAB, tabs) + COMMANDS[command]

		if command == "[" {
			tabs = tabs + 1
		}

		if command == "]" {
			tabs = tabs - 1
		}

		if tabs == 0 {
			return "", NotOpenedCycleError
		}
	}

	println("translated")
	end = end + END
	return end, nil
}


func (bf *Compiler) CompileFile(filename string) error {
	code, err := bf.ReadFile(filename)

	if err != nil {
		return err
	}

	parsed := bf.Parse(code)
	translated, err := bf.Translate(parsed)

	if err != nil {
		return err
	}

	err = bf.WriteFile(strings.Split(filename, ".")[0] + ".go", translated)

	if err != nil {
		return err
	}

	err = bf.CompileCFile(strings.Split(filename, ".")[0])
	return err
}


func (bf *Compiler) CompileCFile(filename string) error {
	cmd := exec.Command("go", "build", filename + ".go")

	if err := cmd.Run(); err != nil {
		return err
	}

	println("compiled")
	return nil
}

func (bf *Compiler) ReadFile(filename string) (string, error) {
	b, err := ioutil.ReadFile(filename)

	if err != nil {
		return "", err
	}

	println("file readed")
	return string(b), nil
}

func (bf *Compiler) WriteFile(filename string, text string) error {
	return ioutil.WriteFile(filename, []byte(text), 0644)
}
