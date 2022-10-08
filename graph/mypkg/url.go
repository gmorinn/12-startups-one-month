package mypkg

import (
	"fmt"
	"io"
	"net/url"
)

type URL string

func (value *URL) UnmarshalGQL(v interface{}) error {
	data, ok := v.(string)
	if !ok {
		return fmt.Errorf("url must be a string")
	}
	u, err := url.ParseRequestURI(data)
	if err != nil {
		return fmt.Errorf("url must be a valid url")
	}
	*value = URL(u.String())
	return nil
}

func (value URL) MarshalGQL(w io.Writer) {
	w.Write([]byte(fmt.Sprintf("\"%s\"", value)))
}
