package generator

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mykytaserdiuk9/httpmock/pkg/models"
	"github.com/mykytaserdiuk9/httpmock/pkg/writer"
	"net/http"
	"net/url"
)

func (r *Runner) addPath(path *models.Path) {
	responser := writer.NewResponser(path.Path)

	for _, endpoint := range path.Endpoints {
		route := r.router.HandleFunc(path.Path, func(writer http.ResponseWriter, request *http.Request) {
			// path parameters validate
			if r.config.ValidatePath {
				pathmap := endpoint.Parameters.PathVars()
				if len(pathmap) != 0 {
					vars := mux.Vars(request)
					if err := ValidatePathVars(pathmap, vars); err != nil {
						responser.Error(writer, err, http.StatusBadRequest)
						return
					}
				}
			}
			if r.config.ValidateQuery {
				querymap := endpoint.Parameters.QueryVars()
				if err := ValidateQueryVars(querymap, request.URL.Query()); err != nil {
					responser.Error(writer, err, http.StatusBadRequest)
					return
				}
			}

			// header
			if r.config.ValidateHeader && endpoint.MayHaveRequest() {
				err := ValidateHeader(endpoint.Request, request)
				responser.Error(writer, err, http.StatusBadRequest)
			}

			// response
			writer.WriteHeader(endpoint.Response.Status)
			writer.Header().Set("Accept-Encoding", endpoint.Response.Type)
			endpoint.Response.Header.WriteTo(writer)

			body, err := json.Marshal(endpoint.Response.Body)
			if err != nil {
				responser.Error(writer, errors.New("error unmarshal response"), http.StatusInternalServerError)
				return
			}
			writer.Write(body)
		})

		route.Methods(endpoint.Method.String())
	}
}

func ValidatePathVars(expected, vars map[string]string) error {
	for k, v := range expected {
		re := vars[k]
		if re == "" {
			return errors.New("there isn`t in path vars variable: " + k)
		}
		if re != v {
			return errors.New("Vars is not equals: " + k)
		}
	}
	return nil
}

func ValidateQueryVars(expected map[string]string, values url.Values) error {
	for k, v := range expected {
		out := values.Get(k)
		if out == "" {
			return errors.New("there isn`t in query vars variable: " + k)
		}
		if out != v {
			return errors.New("Vars is not equals: " + k)
		}
	}
	return nil
}

func ValidateHeader(expected models.Request, request *http.Request) error {
	accept := request.Header["Accept-Encoding"]
	if len(accept) == 0 {
		return errors.New("Accept-Encoding in Request header is empty")
	}
	if accept[0] != expected.Type {
		return fmt.Errorf("Accept-Encoding in Request header is wrong, expected: %s, actual: %s", expected.Type, accept[0])
	}
	err := expected.Header.IsEquals(request)
	if err != nil {
		return err
	}
	return nil
}
