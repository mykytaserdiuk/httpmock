package models

type Request struct {
	Header   Header
	Expected any
	Type     string
}

func (r *Request) IsValid() error {
	// validate type

	// if r.Expected == nil {
	//	return false
	// }
	return nil
}
