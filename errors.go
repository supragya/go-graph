package gograph

import "errors"

var (
	ErrEdgeExists = errors.New("edge exists between given vertices")
)
