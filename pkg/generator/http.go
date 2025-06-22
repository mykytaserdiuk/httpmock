package generator

import (
	"fmt"
	"github.com/mykytaserdiuk9/httpmock/pkg/models"
	"log"
	"os"
)

func (r *Runner) Launch(scheme *models.MockScheme) error {
	if r.config.ValidateScheme {
		err := scheme.IsValid()
		fmt.Print(err)
		os.Exit(1)
	}

	for _, path := range scheme.Paths {
		r.AddPath(path)
	}
	err := r.Run(scheme.Port)
	if err != nil {
		log.Fatalf("Failed to start HTTP server. Err: %s", err.Error())
	}
	return nil
}
