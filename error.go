package legit

import (
	"fmt"
)

// Errors contains one or more validate errors
type Errors []error

// returns string representation of the first validation error encountered
func (e Errors) Error() string {
	if len(e) > 0 {
		return e[0].Error()
	}

	return ""
}

// StructError contains the field name and message of a failed validation
type StructError struct {
	Field   string `json:"field"`
	Message error  `json:"message"`
}

// returns the string representation of the field and failed validation
func (se StructError) Error() string {
	return fmt.Sprintf("%s: %s", se.Field, se.Message)
}

// SliceError contains the index and message of a failed validation
type SliceError struct {
	Index   int   `json:"index"`
	Message error `json:"message"`
}

// returns the string representation of the index and failed validation
func (se SliceError) Error() string {
	return fmt.Sprintf("%d: %s", se.Index, se.Message)
}
