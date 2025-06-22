package generator

import (
	"github.com/mykytaserdiuk9/httpmock/pkg/models"
	"log"
)

func Launch(scheme *models.MockScheme) error {
	runner := NewRunner()
	for _, path := range scheme.Paths {
		runner.AddPath(path)
	}
	err := runner.Run(scheme.Port)
	if err != nil {
		log.Fatalf("Failed to start HTTP server. Err: %s", err.Error())
	}
	return nil
}
