package legit

import (
	"database/sql/driver"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLower(t *testing.T) {
	testString(t, Lower("foo"), Lower("FOO"), errLower)
}

func TestUpper(t *testing.T) {
	testString(t, Upper("FOO"), Upper("foo"), errUpper)
}

func TestNoSpace(t *testing.T) {
	testString(t, NoSpace("foo"), NoSpace(" foo\t bar "), errNoSpace)
}

func TestPrintable(t *testing.T) {
	testString(t, Printable("foo"), Printable("\x00foo"), errPrintable)
}

func TestAlpha(t *testing.T) {
	testString(t, Alpha("foo"), Alpha("abc123"), errAlpha)
}

func TestNumber(t *testing.T) {
	testString(t, Number("1234"), Number("foo"), errNumber)
}

func TestFloat(t *testing.T) {
	testString(t, Float("-1.23"), Float("foo"), errFloat)

	assert.Error(t, Float("").Validate())
	assert.Error(t, Float("1.2.3").Validate())
	assert.Error(t, Float(".1").Validate())
	assert.Error(t, Float("1.").Validate())
}

func TestAlphanumeric(t *testing.T) {
	testString(t, Alphanumeric("abc123"), Alphanumeric(" foo! "), errAlphanumeric)
}

func TestASCII(t *testing.T) {
	testString(t, ASCII("abc123"), ASCII("föö"), errASCII)
}

func TestRequired(t *testing.T) {
	testString(t, Required("foo"), Required(""), errRequired)
}

func testString(t *testing.T, pass, fail Validator, failErr error) {
	assert.NoError(t, pass.Validate())

	err := fail.Validate()
	if assert.NotNil(t, err) {
		assert.Equal(t, failErr, err)
	}
}

func TestStringValue(t *testing.T) {
	tests := []struct {
		Name   string
		Valuer driver.Valuer
		Value  driver.Value
		Error  error
	}{
		{"Lower", Lower("foo"), "foo", nil},
		{"Upper", Upper("foo"), "foo", nil},
		{"NoSpace", NoSpace("foo"), "foo", nil},
		{"Printable", Printable("foo"), "foo", nil},
		{"Alpha", Alpha("foo"), "foo", nil},
		{"Number", Number("foo"), "foo", nil},
		{"Float", Float("foo"), "foo", nil},
		{"Alphanumeric", Alphanumeric("foo"), "foo", nil},
		{"ASCII", ASCII("foo"), "foo", nil},
		{"Required", Required("foo"), "foo", nil},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			v, err := test.Valuer.Value()
			if test.Error == nil {
				if assert.NoError(t, err) {
					assert.Equal(t, test.Value, v)
				}
			} else {
				assert.Equal(t, test.Error, err)
			}
		})
	}
}
