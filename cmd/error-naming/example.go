package main

import (
	"errors"
	"strconv"
)

var (
	errToBig   = errors.New("too big")
	ErrToSmall = errors.New("too small")
	invalid    = errors.New("invalid")
	notError   = 123
)

func doOrError(i int) (string, error) {
	var otherError = errors.New("other")

	if i == 7 {
		return "", otherError
	}

	if i > 10 {
		errMoreThan10 := errors.New("more than 10")
		return "", errMoreThan10
	}

	if i < 5 {
		wrongValue := errors.New("less than 5")
		return "", wrongValue
	}

	return strconv.Itoa(i), nil
}
