package legit

import (
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

func TestAlphanumeric(t *testing.T) {
	testString(t, Alphanumeric("abc123"), Alphanumeric(" foo! "), errAlphanumeric)
}

func TestASCII(t *testing.T) {
	testString(t, ASCII("abc123"), ASCII("föö"), errASCII)
}

func testString(t *testing.T, pass, fail Validator, failErr error) {
	assert.NoError(t, pass.Validate())

	err := fail.Validate()
	if assert.NotNil(t, err) {
		assert.Equal(t, failErr, err)
	}
}
