package models

type Request struct {
	Header   Header
	Expected any
	Type     string
}

func (r *Request) IsValid() bool {
	// validate type

	// if r.Expected == nil {
	//	return false
	// }
	return true
}
