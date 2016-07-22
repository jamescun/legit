package legit

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForm_NewForm(t *testing.T) {
	f := NewForm()
	assert.False(t, f.Legit.Strict)
	assert.Equal(t, Decoders{JSON{}}, f.Decoders)
}

func TestForm_ParseAndValidate(t *testing.T) {
	r := bytes.NewReader([]byte(`"foo"`))

	var body Lower
	err := form.ParseAndValidate(r, "application/json", &body)
	assert.NoError(t, err)
	assert.Equal(t, Lower("foo"), body)
}

func TestForm_ParseAndValidate_noDecoder(t *testing.T) {
	err := form.ParseAndValidate(nil, "application/xml", nil)
	assert.Equal(t, ErrEncoding, err)
}

func TestForm_ParseAndValidate_invalidBody(t *testing.T) {
	r := bytes.NewReader([]byte(`foo`))

	var body string
	err := form.ParseAndValidate(r, "application/json", &body)
	assert.NotNil(t, err)
}

func TestForm_ParseAndValidate_validation(t *testing.T) {
	r := bytes.NewReader([]byte(`"FOO"`))

	var body Lower
	err := form.ParseAndValidate(r, "application/json", &body)
	if assert.NotNil(t, err) {
		assert.Equal(t, errLower, err)
	}
}

func TestForm_ParseRequestAndValidate(t *testing.T) {
	r := &http.Request{
		Body:   ioutil.NopCloser(bytes.NewReader([]byte(`"foo"`))),
		Header: http.Header{"Content-Type": []string{"application/json; charset=utf-8"}},
	}

	var body Lower
	err := form.ParseRequestAndValidate(r, &body)
	assert.NoError(t, err)
	assert.Equal(t, Lower("foo"), body)
}
