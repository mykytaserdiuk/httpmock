package generator

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Runner struct {
	config *Config
	router *mux.Router
}

func NewRunner(config *Config) *Runner {
	return &Runner{config: config, router: mux.NewRouter()}
}

func (r *Runner) Run(port string) error {
	fmt.Printf("\n Mock HTTP Server running on %s", port)
	err := http.ListenAndServe(port, r.router)
	if err != nil {
		return err
	}
	return nil
}
