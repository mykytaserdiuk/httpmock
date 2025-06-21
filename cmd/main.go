package main

import (
	"fmt"
	"github.com/mykytaserdiuk9/httpmock/pkg/generator"
	"gopkg.in/yaml.v3"
	"net/http"

	"github.com/mykytaserdiuk9/httpmock/pkg/models"
)

func main() {
	mock := &models.MockScheme{
		Port: ":8080",
		Paths: []*models.Path{
			{
				Path: "/login/{user_id}",
				Endpoints: models.Endpoints{
					{
						Parameters: []*models.Parameter{{
							In:          "query",
							Placeholder: "user_id",
							Value:       "11",
						}},
						Method: models.MethodGet,
						Response: models.Response{
							Header: map[string][]string{"Agent": {"linux"}},
							Type:   "application/json",
							Status: http.StatusOK,
							Body:   "O KEY",
						},
						Request: models.Request{
							Header:   map[string][]string{"Agent": {"Chrome"}},
							Expected: "55",
							Type:     "application/json",
						},
					},
				}},
		},
	}
	t, _ := yaml.Marshal(mock)
	fmt.Print(string(t))
	err := generator.Launch(mock)
	fmt.Print(err)

}
