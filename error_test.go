package legit

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors_Error(t *testing.T) {
	assert.Equal(t, "foo", Errors{errors.New("foo"), errors.New("bar")}.Error())
	assert.Equal(t, "", Errors{}.Error())
}

func TestStructError_Error(t *testing.T) {
	assert.Equal(t, "foo: bar", StructError{Field: "foo", Message: errors.New("bar")}.Error())
}

func TestSliceError_Error(t *testing.T) {
	assert.Equal(t, "1: bar", SliceError{Index: 1, Message: errors.New("bar")}.Error())
}
