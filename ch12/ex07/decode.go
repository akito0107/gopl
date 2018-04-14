package ex07

import (
	"encoding/json"
	"io"
)

type Decoder struct {
	r       io.Reader
	buf     []byte
	snanned int
}

func NewDecoder(r io.Reader) error {
	json.NewDecoder(r)
	return nil
}

func (d *Decoder) Decode(i interface{}) {
}
