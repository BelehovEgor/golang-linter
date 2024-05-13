package bad

type Foo struct {
	// While these fields technically don't
	// constitute shadowing, grepping for
	// `error` or `string` strings is now
	// ambiguous.
	error  error
	string string
}

func (f Foo) Error() error {
	// `error` and `f.error` are
	// visually similar
	return f.error
}

func (f Foo) String() string {
	// `string` and `f.string` are
	// visually similar
	return f.string
}
