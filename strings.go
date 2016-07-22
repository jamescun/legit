package legit

import (
	"errors"
	"unicode"
)

// Lower validates any string not containing any uppercase characters.
type Lower string

var errLower = errors.New("string is not lowercase")

func (l Lower) Validate() error {
	for _, r := range l {
		if !unicode.IsLower(r) {
			return errLower
		}
	}

	return nil
}

// Upper validates any string not containing any lowercase characters.
type Upper string

var errUpper = errors.New("string is not uppercase")

func (u Upper) Validate() error {
	for _, r := range u {
		if !unicode.IsUpper(r) {
			return errUpper
		}
	}

	return nil
}

// NoSpace validates any string not containing any whitespace characters.
type NoSpace string

var errNoSpace = errors.New("string contains whitespace")

func (ns NoSpace) Validate() error {
	for _, r := range ns {
		if unicode.IsSpace(r) {
			return errNoSpace
		}
	}

	return nil
}

// Printable validates any string not containing any non-printing characters.
type Printable string

var errPrintable = errors.New("string contains non-printing characters")

func (p Printable) Validate() error {
	for _, r := range p {
		if !unicode.IsPrint(r) {
			return errPrintable
		}
	}

	return nil
}

// Alpha validates any string containing only letters.
type Alpha string

var errAlpha = errors.New("string contains non-alpha characters")

func (a Alpha) Validate() error {
	for _, r := range a {
		if !unicode.IsLetter(r) {
			return errAlpha
		}
	}

	return nil
}

// Number validates any string containing only numeric characters.
type Number string

var errNumber = errors.New("string contains non-numeric characters")

func (n Number) Validate() error {
	for _, r := range n {
		if !unicode.IsNumber(r) {
			return errNumber
		}
	}

	return nil
}

// Alphanumeric validates any string containing only letters or numbers.
type Alphanumeric string

var errAlphanumeric = errors.New("string contains non-alphanumeric characters")

func (a Alphanumeric) Validate() error {
	for _, r := range a {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return errAlphanumeric
		}
	}

	return nil
}

// ASCII validates any string containing only ASCII characters.
type ASCII string

var errASCII = errors.New("string contains non-ASCII characters")

func (a ASCII) Validate() error {
	for _, r := range a {
		if r > unicode.MaxASCII {
			return errASCII
		}
	}

	return nil
}
