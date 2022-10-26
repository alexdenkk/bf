package bf

type TokenMap map[string]string

func (tm TokenMap) Parse() {
	for _, key := range KEYS {
		if token, ok := tm[key]; ok {
			if len(token) > 1 {
				tm[key] = key
			}
		} else {
			tm[key] = key
		}
	}
}

var BaseMap = TokenMap{
	"+": "+",
	"-": "-",
	",": ",",
	".": ".",
	"<": "<",
	">": ">",
	"[": "[",
	"]": "]",
}
