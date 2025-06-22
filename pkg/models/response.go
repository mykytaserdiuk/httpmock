package models

import (
	"errors"
	"strconv"
)

type Response struct {
	Type   string
	Status int
	Header Header
	Body   any
}

func (r *Response) IsValid() error {
	// TODO validate type
	// TODO add way to modify validators

	if r.Status < 100 && r.Status > 599 {
		return errors.New("not allowed response status : " + strconv.Itoa(r.Status))
	}
	return nil
}
