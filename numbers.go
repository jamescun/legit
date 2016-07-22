package legit

import (
	"errors"
)

// Positive validates any integer that contains a value above (and including) zero.
type Positive int

var errPositive = errors.New("number is not positive")

func (p Positive) Validate() error {
	if p < 0 {
		return errPositive
	}

	return nil
}

// Negative validates any integer that contains a value below zero.
type Negative int

var errNegative = errors.New("number is not negative")

func (n Negative) Validate() error {
	if n > -1 {
		return errNegative
	}

	return nil
}
