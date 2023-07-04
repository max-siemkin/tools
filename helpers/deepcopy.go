package tools

import (
	"bytes"
	"encoding/gob"
)

func DeepCopy(from, to any) error {
	buff := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	if err := enc.Encode(from); err != nil {
		return err
	}
	return dec.Decode(to)
}
