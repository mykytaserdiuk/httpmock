package generator

import (
	"github.com/mykytaserdiuk9/httpmock/pkg/models"
	"log"
	"net/http"
)

func Launch(scheme *models.MockScheme) error {
	generator := NewGenerator()
	for _, path := range scheme.Paths {
		generator.AddPath(path)
	}
	err := http.ListenAndServe(scheme.Port, generator.router)
	if err != nil {
		log.Fatalf("Failed to start HTTP server. Err: %s", err.Error())
	}
	return nil
}
