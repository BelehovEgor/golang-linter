package bar

import "golang-linter/_tests/guidelines/error/error_bad/foo"

// package bar

func error_bad_usage_example() {
	if err := foo.Open(); err != nil {
		// Can't handle the error.
		panic("unknown error")
	}
}

func error_bad_usage_example2() {

	if err := foo.Open2("testfile.txt"); err != nil {
		// Can't handle the error.
		panic("unknown error")
	}
}
