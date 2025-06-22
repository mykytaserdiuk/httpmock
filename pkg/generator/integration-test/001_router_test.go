package integration_test

import (
	"github.com/golang/mock/gomock"
	"github.com/mykytaserdiuk9/httpmock/pkg/generator"
	generator_mock "github.com/mykytaserdiuk9/httpmock/pkg/generator/mocks"
	"github.com/mykytaserdiuk9/httpmock/pkg/models"
	ta "github.com/mykytaserdiuk9/httpmock/pkg/testassets"
	tu "github.com/mykytaserdiuk9/httpmock/pkg/testassets/testutils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestValidationScheme(t *testing.T) {
	cases := []struct {
		name          string
		input         *models.MockScheme
		expectedError error
	}{{
		name: "OK",
		input: &models.MockScheme{
			Port: ":8080",
			Paths: []*models.Path{{
				Path: "/login/{user_id}",
				Endpoints: models.Endpoints{{
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
				}}},
			}},
		expectedError: nil,
	}, {
		name: "OK_No_Path",
		input: &models.MockScheme{
			Port: "8080",
			Paths: models.Paths{{
				Path:      "",
				Endpoints: nil,
			}},
		},
		expectedError: ta.ErrPathValidationPathEmpty,
	}, {
		name: "OK_No_Endpoints",
		input: &models.MockScheme{
			Port: "8080",
			Paths: models.Paths{{
				Path:      "/api/v1/login",
				Endpoints: make(models.Endpoints, 0),
			}},
		},
		expectedError: ta.ErrValidationZeroEndpoints,
	}, {
		name: "Wrong_Response_Code",
		input: &models.MockScheme{
			Port: ":8080",
			Paths: []*models.Path{{
				Path: "/login/{user_id}",
				Endpoints: models.Endpoints{{
					Parameters: []*models.Parameter{{
						In:          "query",
						Placeholder: "user_id",
						Value:       "11",
					}},
					Method: models.MethodGet,
					Response: models.Response{
						Header: map[string][]string{"Agent": {"linux"}},
						Type:   "application/json",
						Status: 999,
						Body:   "O KEY",
					},
					Request: models.Request{
						Header:   map[string][]string{"Agent": {"Chrome"}},
						Expected: "55",
						Type:     "application/json",
					},
				}}},
			}},
		expectedError: tu.ErrResponseValidationNotAllowedCode(999),
	}, {
		name: "Wrong_Port",
		input: &models.MockScheme{
			Port: "",
			Paths: models.Paths{
				{
					Path:      "",
					Endpoints: nil,
				},
			},
		},
		expectedError: ta.ErrorValidationEmptyPort,
	}, {
		name: "Wrong_Method",
		input: &models.MockScheme{
			Port: ":8080",
			Paths: []*models.Path{{
				Path: "/login/{user_id}",
				Endpoints: models.Endpoints{{
					Parameters: []*models.Parameter{{
						In:          "query",
						Placeholder: "user_id",
						Value:       "11",
					}},
					Method: "WRONG",
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
				}}},
			}},
		expectedError: ta.ErrorValidationWrongMethod,
	}}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			serverMock := generator_mock.NewMockServer(ctrl)
			serverMock.EXPECT().Run(gomock.Any()).Return(nil).MaxTimes(1)

			err := generator.
				NewRunner(ta.ConfigWithValidation, serverMock).
				Launch(c.input)

			assert.Equal(t, c.expectedError, err)
		})
	}
}
