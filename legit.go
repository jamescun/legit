package legit

import (
	"errors"
	"reflect"
)

var (
	// ErrNotStruct is returned when a struct validation method is not given
	// a struct.
	ErrNotStruct = errors.New("object is not a struct")

	// ErrNotSlice is returned when a slice validation method is not given
	// a slice.
	ErrNotSlice = errors.New("object is not a slice")

	// ErrStrict is returned when strict validation mode is enabled and a
	// field does not satisfy the Validator interface.
	ErrStrict = errors.New("field is not a validator")
)

// Validator is a type that can be validated
type Validator interface {
	// returns nil if object is valid
	Validate() error
}

var validator = reflect.TypeOf((*Validator)(nil)).Elem()

var legit = New()

// Legit implements validation of types implementing the Validator interface,
// structs and slices.
type Legit struct {
	// Strict mode requires that all fields in a struct be validatable
	Strict bool
}

// New return a Legit assignment without strict validation
func New() Legit {
	return Legit{
		Strict: false,
	}
}

func Validate(src interface{}) error {
	return legit.Validate(src)
}

func (l Legit) Validate(src interface{}) error {
	// prevent Validation methods being called on nil pointers
	if src == nil {
		return nil
	}

	// skip reflection if src implements custom Validator interface
	if obj, ok := src.(Validator); ok {
		return obj.Validate()
	}

	objv := resolvePointer(reflect.ValueOf(src))
	objt := objv.Type()

	return l.validate(objv, objt)
}

func (l Legit) validate(objv reflect.Value, objt reflect.Type) error {
	// don't attempt to validate pointers (optional fields)
	if objv.Kind() == reflect.Ptr && objv.IsNil() {
		return nil
	}

	if objt.Implements(validator) {
		return objv.Interface().(Validator).Validate()
	}

	switch objv.Kind() {
	case reflect.Struct:
		return l.validateStruct(objv, objt)
	case reflect.Slice:
		return l.validateSlice(objv, objt)
	}

	if l.Strict {
		return ErrStrict
	}

	return nil
}

func ValidateStruct(src interface{}) error {
	return legit.ValidateStruct(src)
}

func (l Legit) ValidateStruct(src interface{}) error {
	objv := resolvePointer(reflect.ValueOf(src))
	if objv.Kind() != reflect.Struct {
		return ErrNotStruct
	}
	objt := objv.Type()

	return l.validateStruct(objv, objt)
}

func (l Legit) validateStruct(objv reflect.Value, objt reflect.Type) error {
	var errors Errors

	for i := 0; i < objt.NumField(); i++ {
		ft := objt.Field(i)

		// see reflect StructField.PkgPath for determining if field is exported
		// TODO: is there a better way to determine if a field is exported?
		if len(ft.PkgPath) < 1 {
			fv := objv.Field(i)

			err := l.validate(fv, fv.Type())
			if err != nil {
				errors = append(errors, StructError{Field: ft.Name, Message: err})
			}
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func ValidateSlice(src interface{}) error {
	return legit.ValidateSlice(src)
}

func (l Legit) ValidateSlice(src interface{}) error {
	objv := resolvePointer(reflect.ValueOf(src))
	if objv.Kind() != reflect.Slice {
		return ErrNotSlice
	}
	objt := objv.Type()

	return l.validateSlice(objv, objt)
}

func (l Legit) validateSlice(objv reflect.Value, objt reflect.Type) error {
	if objv.Len() < 1 {
		return nil
	}

	var errors Errors

	for i := 0; i < objv.Len(); i++ {
		iv := objv.Index(i)
		err := l.validate(iv, iv.Type())
		if err != nil {
			errors = append(errors, SliceError{Index: i, Message: err})
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// return concrete type from arbitrary pointer depth
func resolvePointer(objv reflect.Value) reflect.Value {
	for {
		if objv.Kind() == reflect.Ptr {
			objv = objv.Elem()
		} else {
			break
		}
	}

	return objv
}
