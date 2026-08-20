package bf

import "strings"

func (compiler *Compiler) translate(commands []string) (string, error) {
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
			return "", ErrNotOpenedCycle
		}
	}

	end = end + END
	return end, nil
}
