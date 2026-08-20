package bf

import (
	"errors"
)

var (
	ErrNotOpenedCycle = errors.New("cycle not opened but close operator used")
)
