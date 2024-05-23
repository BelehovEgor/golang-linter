package handleErrorsOnce

import (
	"fmt"
	"log"
)

func exampleHandler1Bad(id int) error {
	u, err := getUser(id)
	if err != nil {
		log.Printf("dddd")
		return fmt.Errorf("get user %q: %w", id, err)
	}
	print(u)
	return nil
}
