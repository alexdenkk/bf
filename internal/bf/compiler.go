package bf

type Compiler struct {
	Tokens []string
}

func NewCompiler() *Compiler {
	return &Compiler{
		Tokens: TOKENS,
	}
}

func (compiler *Compiler) Compile(code string) (string, error) {
	parsed := compiler.parse(code)

	translated, err := compiler.translate(parsed)

	return translated, err
}
