package legit

import (
	"encoding/json"
	"io"
	"strings"
)

// Decoder is used to decode a request body to a type
type Decoder interface {
	// return true if decoder can understand the given MIME type
	Match(mime string) bool

	// return nil if data from reader was unmarshaled into dst
	Decode(r io.Reader, dst interface{}) error
}

// Decoders contains multiple decoders for matching
type Decoders []Decoder

// return the first matching decoder for a MIME type, or nil if no match
func (d Decoders) Match(mime string) Decoder {
	for _, dec := range d {
		if dec.Match(mime) {
			return dec
		}
	}

	return nil
}

// JSON decoder can decode any JSON body with the MIME type "application/json"
type JSON struct{}

func (j JSON) Match(mime string) bool {
	return strings.HasPrefix(mime, "application/json")
}

func (j JSON) Decode(r io.Reader, dst interface{}) error {
	return json.NewDecoder(r).Decode(dst)
}
