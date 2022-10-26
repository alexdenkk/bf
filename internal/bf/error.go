package bf

import (
	"errors"
)

var (
	NotOpenedCycleError = errors.New("cycle not opened but close operator used")
)
