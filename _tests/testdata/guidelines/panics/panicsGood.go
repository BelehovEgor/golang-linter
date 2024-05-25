package panics

import (
	"errors"
	"fmt"
	"os"
)

func runPanicsGood(args []string) error {
	if len(args) == 0 {
		return errors.New("an argument is required")
	}
	return nil
}

func mainPanicsGood() {
	if err := runPanicsGood(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
