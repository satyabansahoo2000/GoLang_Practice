// Custom Error Types
package customErrors

import (
	"fmt"
	"time"
)

type ErrorType struct {
	When time.Time
	what string
}

func (e *ErrorType) Error() string {
	return fmt.Sprintf("Error at %v\nError : %v", e.When, e.what)
}

func run() error {
	return &ErrorType{
		time.Now(),
		"Code not working",
	}
}

func PrintErrors() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}