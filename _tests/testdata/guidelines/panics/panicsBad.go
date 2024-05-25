package panics

import "os"

func runPanicsBad(args []string) {
	if len(args) == 0 {
		panic("an argument is required") // want "If an error occurs, the function must return an error and allow the caller to decide how to handle it."
	}
}

func mainPanicsBad() {
	runPanicsBad(os.Args[1:])
}
