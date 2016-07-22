package legit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositive(t *testing.T) {
	testNumber(t, Positive(100), Positive(-100), errPositive)
}

func TestNegative(t *testing.T) {
	testNumber(t, Negative(-100), Negative(100), errNegative)
}

func testNumber(t *testing.T, pass, fail Validator, failErr error) {
	assert.NoError(t, pass.Validate())

	err := fail.Validate()
	if assert.NotNil(t, err) {
		assert.Equal(t, failErr, err)
	}
}
