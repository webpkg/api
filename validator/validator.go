package validator

import "fmt"

// CreateValidationError create validation error with name and message
func CreateValidationError(name string, message string) error {
	return fmt.Errorf("%s: %s", name, message)
}
