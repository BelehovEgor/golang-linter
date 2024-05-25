package errorNaming

import (
	"errors"
	"fmt"
)

var (
	BrokenLink        = errors.New("link is broken") // want "Global error variable has to start with 'err' or 'Err' prefix"
	CouldNotOpenError = errors.New("could not open") // want "Global error variable has to start with 'err' or 'Err' prefix"
	NotFound          = errors.New("not found")      // want "Global error variable has to start with 'err' or 'Err' prefix"
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
