package legit

import (
	"database/sql/driver"
	"errors"
	"strings"
	"unicode"
)

var errUnsupportedScan = errors.New("unsupported scan")

// Lower validates any string not containing any uppercase characters.
type Lower string

func (l *Lower) Scan(src interface{}) error {
	if s, ok := src.(string); ok {
		*l = Lower(s)
		return nil
	}

	return errUnsupportedScan
}

func (l Lower) Value() (driver.Value, error) {
	return string(l), nil
}

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

func (u *Upper) Scan(src interface{}) error {
	if s, ok := src.(string); ok {
		*u = Upper(s)
		return nil
	}

	return errUnsupportedScan
}

func (u Upper) Value() (driver.Value, error) {
	return string(u), nil
}

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

func (ns *NoSpace) Scan(src interface{}) error {
	if s, ok := src.(string); ok {
		*ns = NoSpace(s)
		return nil
	}

	return errUnsupportedScan
}

func (ns NoSpace) Value() (driver.Value, error) {
	return string(ns), nil
}

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

func (p *Printable) Scan(src interface{}) error {
	if s, ok := src.(string); ok {
		*p = Printable(s)
		return nil
	}

	return errUnsupportedScan
}

func (p Printable) Value() (driver.Value, error) {
	return string(p), nil
}

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

func (a *Alpha) Scan(src interface{}) error {
	if s, ok := src.(string); ok {
		*a = Alpha(s)
		return nil
	}

	return errUnsupportedScan
}

func (a Alpha) Value() (driver.Value, error) {
	return string(a), nil
}

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

func (n *Number) Scan(src interface{}) error {
	if s, ok := src.(string); ok {
		*n = Number(s)
		return nil
	}

	return errUnsupportedScan
}

func (n Number) Value() (driver.Value, error) {
	return string(n), nil
}

var errNumber = errors.New("string contains non-numeric characters")

func (n Number) Validate() error {
	for _, r := range n {
		if !unicode.IsNumber(r) {
			return errNumber
		}
	}

	return nil
}

// Float validates any string containing numbers, including an initial minus
// and a single decimal point.
type Float string

func (f *Float) Scan(src interface{}) error {
	if s, ok := src.(string); ok {
		*f = Float(s)
		return nil
	}

	return errUnsupportedScan
}

func (f Float) Value() (driver.Value, error) {
	return string(f), nil
}

var errFloat = errors.New("float contains non-numeric characters")

func (f Float) Validate() error {
	if len(f) == 0 {
		return errFloat
	}

	s := string(f)
	if s[0] == '-' {
		s = s[1:]
	}

	if strings.Count(s, ".") > 1 {
		return errFloat
	} else if s[0] == '.' || s[len(s)-1] == '.' {
		return errFloat
	}

	i := strings.IndexFunc(s, func(r rune) bool {
		return !((r >= '0' && r <= '9') || r == '.')
	})
	if i > -1 {
		return errFloat
	}

	return nil
}

// Alphanumeric validates any string containing only letters or numbers.
type Alphanumeric string

func (a *Alphanumeric) Scan(src interface{}) error {
	if s, ok := src.(string); ok {
		*a = Alphanumeric(s)
		return nil
	}

	return errUnsupportedScan
}

func (a Alphanumeric) Value() (driver.Value, error) {
	return string(a), nil
}

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

func (a *ASCII) Scan(src interface{}) error {
	if s, ok := src.(string); ok {
		*a = ASCII(s)
		return nil
	}

	return errUnsupportedScan
}

func (a ASCII) Value() (driver.Value, error) {
	return string(a), nil
}

var errASCII = errors.New("string contains non-ASCII characters")

func (a ASCII) Validate() error {
	for _, r := range a {
		if r > unicode.MaxASCII {
			return errASCII
		}
	}

	return nil
}

// Required validates any string that is not empty
type Required string

func (r *Required) Scan(src interface{}) error {
	if s, ok := src.(string); ok {
		*r = Required(s)
		return nil
	}

	return errUnsupportedScan
}

func (r Required) Value() (driver.Value, error) {
	return string(r), nil
}

var errRequired = errors.New("string is required")

func (r Required) Validate() error {
	if len(r) < 1 {
		return errRequired
	}

	return nil
}
