package db

import (
	"errors"
)

var (
	NoResultErr = errors.New("db returned empty result")
)
