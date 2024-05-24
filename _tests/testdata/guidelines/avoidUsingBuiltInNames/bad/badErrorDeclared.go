package bad

// should be detected with default linters
var error string

// `error` shadows the builtin

// or

func handleErrorMessage(error string) {
	// `error` shadows the builtin
}
