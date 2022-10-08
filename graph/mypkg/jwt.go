package mypkg

import (
	"fmt"
	"io"
	"regexp"
)

type JWT string

func (value *JWT) UnmarshalGQL(v interface{}) error {
	data, ok := v.(string)
	if !ok {
		return fmt.Errorf("jwt must be a string")
	}
	tokenRegex := regexp.MustCompile(`^[a-zA-Z0-9\-_]+?\.[a-zA-Z0-9\-_]+?\.([a-zA-Z0-9\-_]+)?$`)
	if tokenRegex.MatchString(data) {
		*value = JWT(data)
	} else {
		return fmt.Errorf("jwt must be a valid jwt token")
	}
	return nil
}

func (value JWT) MarshalGQL(w io.Writer) {
	w.Write([]byte(fmt.Sprintf("\"%s\"", value)))
}
