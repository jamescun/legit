package legit

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLegit_New(t *testing.T) {
	v := New()
	assert.False(t, v.Strict)
}

func TestValidate(t *testing.T) {
	err := Validate(struct {
		Name Lower
	}{"FOO"})
	if assert.NotNil(t, err) {
		assert.Equal(t, Errors{StructError{Field: "Name", Message: errLower}}, err)
	}
}

func TestLegit_Validate(t *testing.T) {
	err := legit.Validate([]Lower{"FOO"})
	if assert.NotNil(t, err) {
		assert.Equal(t, Errors{SliceError{Index: 0, Message: errLower}}, err)
	}
}

func TestLegit_Validate_customValidator(t *testing.T) {
	err := legit.Validate(Lower("foo"))
	assert.NoError(t, err)

	err = legit.Validate(Lower("FOO"))
	if assert.NotNil(t, err) {
		assert.Equal(t, errLower, err)
	}
}

func TestLegit_validate_struct(t *testing.T) {
	err := legit.validate(reflected(struct {
		Name Lower
	}{"FOO"}))
	if assert.NotNil(t, err) {
		assert.Equal(t, Errors{StructError{Field: "Name", Message: errLower}}, err)
	}
}

func TestLegit_validate_slice(t *testing.T) {
	err := legit.validate(reflected([]Lower{"FOO"}))
	if assert.NotNil(t, err) {
		assert.Equal(t, Errors{SliceError{Index: 0, Message: errLower}}, err)
	}
}

func TestLegit_validate_customValidator(t *testing.T) {
	err := legit.validate(reflected(Lower("foo")))
	assert.NoError(t, err)

	err = legit.validate(reflected(Lower("FOO")))
	if assert.NotNil(t, err) {
		assert.Equal(t, errLower, err)
	}
}

func TestLegit_validate_strict(t *testing.T) {
	l := Legit{Strict: true}
	err := l.validate(reflected("foo"))
	assert.Equal(t, ErrStrict, err)
}

func TestLegit_validate_unknown(t *testing.T) {
	err := legit.validate(reflected("foo"))
	assert.NoError(t, err)
}

func TestValidateStruct(t *testing.T) {
	err := ValidateStruct(struct {
		Name Lower
	}{"FOO"})
	if assert.NotNil(t, err) {
		assert.Equal(t, Errors{StructError{Field: "Name", Message: errLower}}, err)
	}

}

func TestLegit_ValidateStruct(t *testing.T) {
	err := legit.ValidateStruct(Lower("foo"))
	assert.Equal(t, ErrNotStruct, err)

	err = legit.ValidateStruct(struct {
		Name Lower
	}{"FOO"})
	if assert.NotNil(t, err) {
		assert.Equal(t, Errors{StructError{Field: "Name", Message: errLower}}, err)
	}
}

func TestLegit_validateStruct(t *testing.T) {
	err := legit.validateStruct(reflected(struct {
		Name Lower
		name Lower // unexported fields should NOT be validated
	}{"foo", "FOO"}))
	assert.NoError(t, err)

	err = legit.validateStruct(reflected(struct {
		First Lower
		Last  Lower
	}{"foo", "FOO"}))
	if assert.NotNil(t, err) {
		assert.Equal(t, Errors{StructError{Field: "Last", Message: errLower}}, err)
	}
}

func TestValidateSlice(t *testing.T) {
	err := ValidateSlice([]Lower{"FOO"})
	if assert.NotNil(t, err) {
		assert.Equal(t, Errors{SliceError{Index: 0, Message: errLower}}, err)
	}
}

func TestLegit_ValidateSlice(t *testing.T) {
	err := legit.ValidateSlice(Lower("foo"))
	assert.Equal(t, ErrNotSlice, err)

	err = legit.ValidateSlice([]Lower{"FOO"})
	if assert.NotNil(t, err) {
		assert.Equal(t, Errors{SliceError{Index: 0, Message: errLower}}, err)
	}
}

func TestLegit_validateSlice(t *testing.T) {
	err := legit.validateSlice(reflected([]Lower{}))
	assert.NoError(t, err)

	err = legit.validateSlice(reflected([]Lower{"foo"}))
	assert.NoError(t, err)

	err = legit.validateSlice(reflected([]Lower{"foo", "FOO"}))
	if assert.NotNil(t, err) {
		assert.Equal(t, Errors{SliceError{Index: 1, Message: errLower}}, err)
	}
}

func TestResolvePointer(t *testing.T) {
	v := resolvePointer(reflect.ValueOf(&struct{ Foo string }{"bar"}))
	assert.Equal(t, reflect.Struct, v.Kind())
}

func reflected(src interface{}) (objv reflect.Value, objt reflect.Type) {
	objv = reflect.ValueOf(src)
	objt = objv.Type()
	return
}
