package schema_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mykytaserdiuk9/httpmock/pkg/models"
	"github.com/mykytaserdiuk9/httpmock/pkg/schema"
)

func TestIsURL(t *testing.T) {
	testCases := []struct {
		name     string
		source   string
		expected bool
	}{
		{
			name:     "Valid HTTP URL",
			source:   "http://google.com",
			expected: true,
		}, {
			name:     "Valid HTTPS URL",
			source:   "https://google.com",
			expected: true,
		}, {
			name:     "Local File Path",
			source:   "./api/example.yaml",
			expected: false,
		}, {
			name:     "Valid HTTP URL, with trailing slash",
			source:   "https:///google.com",
			expected: true,
		}, {
			name:     "Valid HTTP URL, without one slash",
			source:   "https:/google.com",
			expected: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ok := schema.IsURL(tc.source)
			assert.Equal(t, tc.expected, ok)
		})
	}
}

func TestDownloadSchema(t *testing.T) {
	testCases := []struct {
		name          string
		source        string
		expectedBody  string
		expectedError string
	}{
		{
			name:   "OK",
			source: "https://jsonplaceholder.typicode.com/todos/1",
			expectedBody: `{
  "userId": 1,
  "id": 1,
  "title": "delectus aut autem",
  "completed": false
}`,
		}, {
			name:          "Local File Path",
			source:        "./local/path/to/file.yaml",
			expectedError: "unsupported protocol scheme",
		}, {
			name:          "NotFound",
			source:        "http://hfhhfhf.com",
			expectedError: "no such host",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, err := schema.DownloadSchema(tc.source)
			if err != nil {
				assert.ErrorContains(t, err, tc.expectedError)
			}
			assert.Equal(t, tc.expectedBody, string(body))
		})
	}
}

func TestGet(t *testing.T) {
	testCases := []struct {
		name         string
		source       string
		expectedBody *models.MockScheme
	}{
		{
			name:   "OK",
			source: "https://raw.githubusercontent.com/mykytaserdiuk/httpmock/refs/heads/main/api/example.yaml",
			expectedBody: &models.MockScheme{
				Port: ":8080",
				Paths: []*models.Path{{
					Path: "/login/{user_id}",
					Endpoints: models.Endpoints{{
						Parameters: []*models.Parameter{{
							In:          "path",
							Placeholder: "user_id",
							Value:       "11",
						}},
						Method: models.MethodGet,
						Response: models.Response{
							Header: map[string][]string{"AGENT": {"linux"}},
							Type:   "application/json",
							Status: http.StatusOK,
							Body:   "O KEY",
						},
						Request: models.Request{
							Header:   map[string][]string{"AGENT": {"Chrome"}},
							Expected: "55",
							Type:     "application/json",
						},
					}}},
				}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, err := schema.Get(tc.source)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedBody, body)
		})
	}
}
