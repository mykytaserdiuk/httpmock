package generator

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Runner struct {
	router *mux.Router
}

func NewRunner() *Runner {
	return &Runner{router: mux.NewRouter()}
}

func (r *Runner) Run(port string) error {
	err := http.ListenAndServe(port, r.router)
	if err != nil {
		return err
	}
	return nil
}
