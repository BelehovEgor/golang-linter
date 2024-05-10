package panics

import "os"

func runPanicsBad(args []string) {
	if len(args) == 0 {
		panic("an argument is required")
	}
}

func mainPanicsBad() {
	runPanicsBad(os.Args[1:])
}
