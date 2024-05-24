package good

type Foo struct {
	// ...
}

var _defaultFoo1 = Foo{
	// ...
}

// or, better, for testability:

var _defaultFoo2 = defaultFoo()

func defaultFoo() Foo {
	return Foo{
		// ...
	}
}
