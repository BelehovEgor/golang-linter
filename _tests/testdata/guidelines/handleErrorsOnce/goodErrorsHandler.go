package handleErrorsOnce

import (
	"errors"
	"fmt"
	"log"
)

type User struct {
}

func getUser(id int) (User, error) {
	return User{}, nil
}

var SomeRandomError = errors.New("")

func exampleHandler1Good(id int) error {
	u, err := getUser(id)
	if err != nil {
		if errors.Is(err, SomeRandomError) {
			// Some metrics should not
			// break the application.
			log.Printf("Could not emit metrics: %v", err)
		} else {
			// No log when return error
			return fmt.Errorf("get user %q: %w", id, err)
		}
	}
	print(u)
	return nil
}
