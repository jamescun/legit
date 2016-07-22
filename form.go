package legit

import (
	"errors"
	"io"
	"net/http"
)

var (
	// ErrEncoding is returned when a matching decoder is not found for
	// the encoding.
	ErrEncoding = errors.New("unknown encoding")
)

// Form implements the decoding and validation of user data from readers
// and HTTP requests
type Form struct {
	Legit    Legit
	Decoders Decoders
}

var form = NewForm()

// NewForm returns a Form assignment with the default Legit configuration and
// a JSON decoder
func NewForm() Form {
	return Form{
		Legit:    New(),
		Decoders: Decoders{JSON{}},
	}
}

// ParseAndValidate first decodes a reader using the first decoder matching the
// given mime type, then applies validation to the collected input
func (f Form) ParseAndValidate(r io.Reader, mime string, dst interface{}) error {
	dec := f.Decoders.Match(mime)
	if dec == nil {
		return ErrEncoding
	}

	err := dec.Decode(r, dst)
	if err != nil {
		return err
	}

	err = f.Legit.Validate(dst)
	if err != nil {
		return err
	}

	return nil
}

// ParseRequestAndValidate is the same as ParseAndValidate accepting a HTTP
// request for the reader and using the "Content-Type" header for the MIME type
func (f Form) ParseRequestAndValidate(r *http.Request, dst interface{}) error {
	return f.ParseAndValidate(r.Body, r.Header.Get("Content-Type"), dst)
}
