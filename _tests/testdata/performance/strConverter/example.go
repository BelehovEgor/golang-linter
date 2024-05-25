package strConverter

import (
	"fmt"
	"log"
)

func myLog(format string, args ...interface{}) {
	log.Printf(format, args...)
	fmt.Sprint(123) // want "Use 'strconv' instead 'fmt' for converting type to/from string"
}
