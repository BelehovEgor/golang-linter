package foo

import (
	"errors"
	"fmt"
)

// package foo

func Open() error {
	return errors.New("could not open")
}

// package foo

func Open2(file string) error {
	return fmt.Errorf("file %q not found", file)
}
