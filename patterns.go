package legit

import (
	"errors"
	"regexp"
)

// Email validates any string matching a RFC 5322 email address
type Email string

// RFC 5322 "official" email regexp
var expEmail = regexp.MustCompile(`(?:[a-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])`)

var errEmail = errors.New("invalid email")

func (e Email) Validate() error {
	if !expEmail.MatchString(string(e)) {
		return errEmail
	}

	return nil
}

// CreditCard validates any string matching a credit card number
type CreditCard string

var expCreditCard = regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`)

var errCreditCard = errors.New("invalid credit card")

func (c CreditCard) Validate() error {
	if !expCreditCard.MatchString(string(c)) {
		return errCreditCard
	}

	return nil
}

// UUID validates any string matching a UUID of any version
type UUID string

var expUUID = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)

var errUUID = errors.New("invalid uuid")

func (u UUID) Validate() error {
	if !expUUID.MatchString(string(u)) {
		return errUUID
	}

	return nil
}

// UUID3 validates any string matching a version 3 UUID
type UUID3 string

var expUUID3 = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$`)

var errUUID3 = errors.New("invalid uuid3")

func (u UUID3) Validate() error {
	if !expUUID3.MatchString(string(u)) {
		return errUUID3
	}

	return nil
}

// UUID4 validates any string matching a version 4 UUID
type UUID4 string

var expUUID4 = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)

var errUUID4 = errors.New("invalid uuid4")

func (u UUID4) Validate() error {
	if !expUUID4.MatchString(string(u)) {
		return errUUID4
	}

	return nil
}

// UUID5 validates any string matching a version 5 UUID
type UUID5 string

var expUUID5 = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)

var errUUID5 = errors.New("invalid uuid5")

func (u UUID5) Validate() error {
	if !expUUID5.MatchString(string(u)) {
		return errUUID5
	}

	return nil
}
