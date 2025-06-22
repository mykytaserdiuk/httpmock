package generator

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/mykytaserdiuk9/httpmock/pkg/models"
	"net/http"
	"net/url"
)

func (r *Runner) AddPath(path *models.Path) {
	for _, endpoint := range path.Endpoints {
		route := r.router.HandleFunc(path.Path, func(writer http.ResponseWriter, request *http.Request) {
			// path parameters validate
			pathmap := endpoint.Parameters.PathVars()
			if len(pathmap) != 0 {
				vars := mux.Vars(request)
				if err := ValidatePathVars(pathmap, vars); err != nil {
					// TODO Write to runner
					http.Error(writer, err.Error(), http.StatusBadRequest)
					return
				}
			}

			querymap := endpoint.Parameters.QueryVars()
			if err := ValidateQueryVars(querymap, request.URL.Query()); err != nil {
				http.Error(writer, err.Error(), http.StatusBadRequest)
				return
			}

			// request
			if endpoint.MayHaveRequest() {
				accept := request.Header["Accept-Encoding"][0]
				if accept != endpoint.Request.Type {
					http.Error(writer, "Accept-Encoding in request is wrong", http.StatusBadRequest)
					return
				}
				err := endpoint.Request.Header.IsEquals(request)
				if err != nil {
					http.Error(writer, err.Error(), http.StatusBadRequest)
					return
				}
			}

			// response
			writer.WriteHeader(endpoint.Response.Status)
			writer.Header().Set("Accept-Encoding", endpoint.Response.Type)
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

func ValidatePathVars(expected, vars map[string]string) error {
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

func ValidateQueryVars(expected map[string]string, values url.Values) error {
	for k, v := range expected {
		out := values.Get(k)
		if out == "" {
			return errors.New("there isn`t in vars variable: " + k)
		}
		if out != v {
			return errors.New("Vars is not equals: " + k)
		}
	}
	return nil
}
