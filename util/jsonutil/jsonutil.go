// Package jsonutil provides JSON unmarshaling utilities for types.
package jsonutil

import (
	"encoding/json"
	"strconv"
)

// Int is used to parse strings as Go integers.
type Int struct {
	value uint64
}

// UnmarshalJSON implements json.Unmarshaler.
func (i *Int) UnmarshalJSON(data []byte) error {
	var (
		s   string
		err error
	)
	if err = json.Unmarshal(data, &s); err != nil {
		return err
	}
	i.value, err = strconv.ParseUint(s, 10, 64)
	return err
}

// Int64 creates an int64 from Int.
func (i *Int) Int64() int64 {
	return int64(i.value)
}
