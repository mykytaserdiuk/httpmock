package models

import (
	"errors"
	"net/http"
	"reflect"
	"sort"
)

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

func (h Header) IsEquals(r *http.Request) error {
	actual := r.Header

	for key, expectedVals := range h {
		actualVals, ok := actual[key]
		if !ok {
			return errors.New("missing header key: '" + key + "' in actual response")
		}

		expected := append([]string{}, expectedVals...)
		actualCopy := append([]string{}, actualVals...)

		sort.Strings(expected)
		sort.Strings(actualCopy)

		if !reflect.DeepEqual(expected, actualCopy) {
			return errors.New("header values for key '" + key + "' do not match")
		}
	}

	return nil
}
