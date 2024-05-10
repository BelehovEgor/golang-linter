package foo

import (
	"errors"
	"fmt"
)

// package foo

var ErrCouldNotOpen = errors.New("could not open")

func Open() error {
	return ErrCouldNotOpen
}

// package foo

type NotFoundError struct {
	File string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("file %q not found", e.File)
}

func Open2(file string) error {
	return &NotFoundError{File: file}
}

// package bar
