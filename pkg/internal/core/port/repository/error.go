package repository

import (
	"errors"
)

type Code int

var (
	NotFound = errors.New("could not find")
)
