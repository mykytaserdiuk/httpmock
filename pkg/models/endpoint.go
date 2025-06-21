package models

import (
	"strings"
)

const (
	MethodGet     HTTPMethod = "GET"
	MethodHead    HTTPMethod = "HEAD"
	MethodPost    HTTPMethod = "POST"
	MethodPut     HTTPMethod = "PUT"
	MethodPatch   HTTPMethod = "PATCH"
	MethodDelete  HTTPMethod = "DELETE"
	MethodConnect HTTPMethod = "CONNECT"
	MethodOptions HTTPMethod = "OPTIONS"
	MethodTrace   HTTPMethod = "TRACE"
)

type HTTPMethod string

func (m HTTPMethod) IsValid() bool {
	switch strings.ToUpper(string(m)) {
	case string(MethodGet), string(MethodHead), string(MethodPost),
		string(MethodPut), string(MethodPatch), string(MethodDelete),
		string(MethodConnect), string(MethodOptions), string(MethodTrace):
		return true
	default:
		return false
	}
}
func (m HTTPMethod) IsUpdating() bool {
	switch strings.ToUpper(string(m)) {
	case string(MethodPost),
		string(MethodPut), string(MethodPatch), string(MethodDelete),
		string(MethodConnect):
		return true
	default:
		return false
	}
}
func (m HTTPMethod) String() string {
	return string(m)
}

type Endpoints []*Endpoint
type Endpoint struct {
	Parameters Parameters
	Method     HTTPMethod
	Request    Request
	Response   Response
}

func (e *Endpoint) IsValid() bool {
	if ok := e.Method.IsValid(); !ok {
		return false
	}
	if e.Method.IsUpdating() {
		if ok := e.Request.IsValid(); !ok {
			return false
		}
	} else {
		// TODO add checking response value if "GETTING"
	}

	if ok := e.Response.IsValid(); !ok {
		return false
	}
	return true
}
func (e *Endpoint) MayHaveRequest() bool {
	return !e.Method.IsUpdating()
}

type Parameters []*Parameter
type Parameter struct {
	In          string
	Placeholder string
	Value       string
}

func (p Parameters) Query() map[string]string {
	resp := make(map[string]string, 0)
	for _, par := range p {
		if par.In == "query" {
			resp[par.Placeholder] = par.Value
		}
	}
	return resp
}
