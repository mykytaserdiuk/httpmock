package generator

import (
	"net/http"

	"github.com/mykytaserdiuk9/httpmock/pkg/models"
)

type Server interface {
	Run(handler http.Handler) error
}

type Launcher interface {
	Launch(scheme *models.MockScheme) error
}
