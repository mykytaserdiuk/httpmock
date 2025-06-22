package generator

import (
	"fmt"
	"log"

	"github.com/gorilla/mux"

	"github.com/mykytaserdiuk9/httpmock/pkg/models"
)

type Runner struct {
	config *Config
	router *mux.Router
	server Server
}

func NewRunner(config *Config, server Server) *Runner {
	return &Runner{
		config: config,
		router: mux.NewRouter(),
		server: server,
	}
}

func (r *Runner) Launch(scheme *models.MockScheme) error {
	if r.config.ValidateScheme {
		if err := scheme.IsValid(); err != nil {
			return fmt.Errorf("error validation scheme : %s", err)
		}
	}

	for _, path := range scheme.Paths {
		r.addPath(path)
	}
	err := r.server.Run(r.router)
	if err != nil {
		log.Fatalf("Failed to start HTTP server. Err: %s", err.Error())
	}
	return nil
}
