package errorNaming

import (
	"errors"
	"fmt"
)

var (
	// The following two errors are exported
	// so that users of this package can match them
	// with errors.Is.

	ErrBrokenLink   = errors.New("link is broken")
	ErrCouldNotOpen = errors.New("could not open")

	// This error is not exported because
	// we don't want to make it part of our public API.
	// We may still use it inside the package
	// with errors.Is.

	errNotFound = errors.New("not found")
)

// Similarly, this error is exported
// so that users of this package can match it
// with errors.As.

type NotFoundError struct {
	File string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("file %q not found", e.File)
}

// And this error is not exported because
// we don't want to make it part of the public API.
// We can still use it inside the package
// with errors.As.

type resolveError struct {
	Path string
}

func (e *resolveError) Error() string {
	return fmt.Sprintf("resolve %q", e.Path)
}
