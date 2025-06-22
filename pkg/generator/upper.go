package generator

import (
	"fmt"
	"net/http"
)

type HTTPUpper struct {
	port string
}

func NewUpper(port string) *HTTPUpper {
	return &HTTPUpper{
		port: port,
	}
}

func (s *HTTPUpper) Run(handler http.Handler) error {
	fmt.Printf("\n Mock HTTP HTTPUpper running on %s", s.port)
	err := http.ListenAndServe(s.port, handler)
	if err != nil {
		return err
	}
	return nil
}
