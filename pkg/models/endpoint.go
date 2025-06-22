package models

import (
	"errors"
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

func (m HTTPMethod) IsValid() error {
	switch strings.ToUpper(string(m)) {
	case string(MethodGet), string(MethodHead), string(MethodPost),
		string(MethodPut), string(MethodPatch), string(MethodDelete),
		string(MethodConnect), string(MethodOptions), string(MethodTrace):
		return nil
	default:
		return errors.New("not allowed method : " + m.String())
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

func (es Endpoints) IsValid() error {
	for _, e := range es {
		return e.IsValid()
	}
	return nil
}
func (e *Endpoint) IsValid() error {
	if err := e.Method.IsValid(); err != nil {
		return err
	}
	if e.Method.IsUpdating() {
		if err := e.Request.IsValid(); err != nil {
			return err
		}
	} else {
		// TODO add checking response value if "GETTING"
	}

	if err := e.Response.IsValid(); err != nil {
		return errors.New("response is not valid : " + err.Error())
	}
	return nil
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

func (p Parameters) PathVars() map[string]string {
	return p.getFrom("path")
}

func (p Parameters) QueryVars() map[string]string {
	return p.getFrom("query")
}

func (p Parameters) getFrom(in string) map[string]string {
	resp := make(map[string]string, 0)
	for _, par := range p {
		if strings.ToLower(par.In) == in {
			resp[par.Placeholder] = par.Value
		}
	}
	return resp
}
