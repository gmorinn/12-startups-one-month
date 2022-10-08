package mypkg

import (
	"fmt"
	"io"

	"github.com/google/uuid"
)

type UUID string

func (value *UUID) UnmarshalGQL(v interface{}) error {
	data, ok := v.(string)
	if !ok {
		return fmt.Errorf("id must be a string")
	}
	id, err := uuid.Parse(data)
	if err != nil {
		return fmt.Errorf("id must be a valid UUID")
	}
	*value = UUID(id.String())
	return nil
}

func (value UUID) MarshalGQL(w io.Writer) {
	w.Write([]byte(fmt.Sprintf("\"%s\"", value)))
}
