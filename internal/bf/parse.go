package bf

import "slices"

func (compiler *Compiler) parse(commands string) []string {
	parsed := []string{}

	for _, command := range commands {
		if slices.Contains(compiler.Tokens, string(command)) {
			parsed = append(parsed, string(command))
		}
	}

	return parsed
}
