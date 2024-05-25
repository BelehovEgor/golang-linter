package panics

func do() {
	panic("some panic") // want "If an error occurs, the function must return an error and allow the caller to decide how to handle it."
}
