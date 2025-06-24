package schema

import (
	"github.com/mykytaserdiuk9/httpmock/pkg/cfg"
	"github.com/mykytaserdiuk9/httpmock/pkg/models"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func IsURL(source string) bool {
	source = strings.TrimSpace(source)
	if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") {
		_, err := url.Parse(source)
		if err != nil {
			return false
		}
		return true
	}

	return false
}

func DownloadSchema(source string) ([]byte, error) {
	resp, err := http.Get(source)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func Get(source string) (*models.MockScheme, error) {
	if IsURL(source) {
		b, err := DownloadSchema(source)
		if err != nil {
			return nil, err
		}
		var mock models.MockScheme
		err = cfg.UnmarshalYAMLRaw(b, &mock)
		if err != nil {
			return nil, err
		}
		return &mock, nil
	}

	// If the source is not a URL, treat it as a local file path
	var mock models.MockScheme
	err := cfg.UnmarshalYAML(source, &mock)
	if err != nil {
		return nil, err
	}
	return &mock, nil

}
