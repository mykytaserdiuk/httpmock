package models

type Response struct {
	Type   string
	Status int
	Header Header
	Body   any
}

func (r *Response) IsValid() bool {
	// TODO validate type
	// TODO add way to modify validators

	if r.Status < 100 && r.Status > 599 {
		return false
	}
	return true
}
