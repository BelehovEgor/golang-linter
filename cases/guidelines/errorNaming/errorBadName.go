package errorNaming

import (
	"errors"
	"fmt"
)

var (
	BrokenLink        = errors.New("link is broken")
	CouldNotOpenError = errors.New("could not open")

	NotFound = errors.New("not found")
)

type ErrorNotFound struct {
	File string
}

func (e *ErrorNotFound) Error2() string {
	return fmt.Sprintf("file %q not found", e.File)
}

type resolveError2 struct {
	Path string
}

func (e *resolveError) Error2() string {
	return fmt.Sprintf("resolve %q", e.Path)
}
