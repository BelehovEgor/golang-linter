package good

var errorMessage string

// `error` refers to the builtin

// or

func handleErrorMessage(msg string) {
	// `error` refers to the builtin
}

type Foo struct {
	// `error` and `string` strings are
	// now unambiguous.
	err error
	str string
}

func (f Foo) Error() error {
	return f.err
}

func (f Foo) String() string {
	return f.str
}
