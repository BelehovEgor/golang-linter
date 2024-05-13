package bad

type Foo struct {
	// ...
}

var _defaultFoo Foo

func init() {
	_defaultFoo = Foo{
		// ...
	}
}
