package legit

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecoders_Match(t *testing.T) {
	d := Decoders{JSON{}}
	assert.Equal(t, JSON{}, d.Match("application/json; charset=utf-8"))
	assert.Equal(t, nil, d.Match("application/xml"))
}

func TestJSON_Match(t *testing.T) {
	j := JSON{}
	assert.True(t, j.Match("application/json"))
	assert.True(t, j.Match("application/json; charset=utf-8"))
	assert.False(t, j.Match("application/xml"))
}

func TestJSON_Decode(t *testing.T) {
	j := JSON{}
	r := bytes.NewReader([]byte(`"hello world"`))

	var body string
	err := j.Decode(r, &body)
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", body)
	}
}
