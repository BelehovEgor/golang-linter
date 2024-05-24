package bar

import (
	"errors"
	"golang-linter/_tests/guidelines/error/error_good/foo"
)

func some_func() {
	// package bar

	if err := foo.Open(); err != nil {
		if errors.Is(err, foo.ErrCouldNotOpen) {
			// handle the error
		} else {
			panic("unknown error")
		}
	}

}

func some_func2() {

	if err := foo.Open2("testfile.txt"); err != nil {
		var notFound *foo.NotFoundError
		if errors.As(err, &notFound) {
			// handle the error
		} else {
			panic("unknown error")
		}
	}
}
