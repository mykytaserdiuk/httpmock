package generator

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/mykytaserdiuk9/httpmock/pkg/models"
	"net/http"
)

func (g *Generator) AddPath(path *models.Path) {
	for _, endpoint := range path.Endpoints {
		route := g.router.HandleFunc(path.Path, func(writer http.ResponseWriter, request *http.Request) {
			qmap := endpoint.Parameters.Query()
			if len(qmap) != 0 {
				vars := mux.Vars(request)
				if err := ValidateVars(qmap, vars); err != nil {
					// TODO Write to generator
					http.Error(writer, err.Error(), http.StatusBadRequest)
					return
				}
			}
			writer.WriteHeader(endpoint.Response.Status)
			endpoint.Response.Header.WriteTo(writer)
			body, err := json.Marshal(endpoint.Response.Body)
			if err != nil {
				// TODO add pkg to write good response for understand
				http.Error(writer, "error unmarshal response", http.StatusInternalServerError)
				return
			}
			writer.Write(body)
		})

		route.Methods(endpoint.Method.String())
	}
}

func ValidateVars(expected, vars map[string]string) error {
	for k, v := range expected {
		re := vars[k]
		if re == "" {
			return errors.New("there isn`t in vars variable: " + k)
		}
		if re != v {
			return errors.New("Vars is not equals: " + k)
		}
	}
	return nil
}
