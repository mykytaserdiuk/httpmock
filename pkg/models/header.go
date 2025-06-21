package models

import "net/http"

// Header can have expected properties
// TODO or some logic
type Header map[string][]string

func (h Header) WriteTo(w http.ResponseWriter) {
	for key, val := range h {
		for _, v := range val {
			w.Header().Add(key, v)
		}
	}
}
