package data

import (
	"encoding/json"
	"io"
)

// ToJson serializes the given interface into a string base JSON format
func ToJson(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

// FromJson deserializes the object from JSON string
func FromJson(i interface{}, r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(i)
}
