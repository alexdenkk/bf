package bf

const BEGIN = `// Generated by brainfuck compiler
// @alexdenkk

package main

import (
	"bufio"
	"os"
)

func main() {

	memory := []rune{}
	cursor := 0

	for i := 0; i < 30000; i++ {
		memory = append(memory, 0)
	}
`

const END = `
}

func input() rune {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return rune([]byte(input)[0])
}
`

const TAB = "    "

var COMMANDS = map[string]string{
	"+": "memory[cursor]++",
	"-": "memory[cursor]--",
	",": "memory[cursor] = input()",
	".": "println(string(memory[cursor]))",
	"<": "cursor--",
	">": "cursor++",
	"[": "for memory[cursor] != 0 {",
	"]": "}",
}