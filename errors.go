package jalaali

import "fmt"

type ErrorNilReference struct{}

type ErrorInvalidYear struct {
	year int
}

func (e *ErrorNilReference) Error() string {
	return "jalaali: reference is nil"
}

func (e *ErrorInvalidYear) Error() string {
	return fmt.Sprintf("jalaali: %v is invalid year", e.year)
}
