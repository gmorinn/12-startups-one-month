package mypkg

import (
	"fmt"
	"io"
	"regexp"
)

type Email string

func (value *Email) UnmarshalGQL(v interface{}) error {
	data, ok := v.(string)
	if !ok {
		return fmt.Errorf("email must be a string")
	}
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if emailRegex.MatchString(data) {
		*value = Email(data)
	} else {
		return fmt.Errorf("Email must be a valid email")
	}
	return nil
}

func (value Email) MarshalGQL(w io.Writer) {
	w.Write([]byte(fmt.Sprintf("\"%s\"", value)))
}
