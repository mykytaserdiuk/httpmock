package testassets

import (
	"errors"
	"github.com/mykytaserdiuk9/httpmock/pkg/generator"
)

var (
	ErrPathValidationPathEmpty = errors.New("error validation scheme : one of path is empty")
	ErrValidationZeroEndpoints = errors.New("error validation scheme : zero endpoints one of path")
	ErrorValidationEmptyPort   = errors.New("error validation scheme : port is empty")
	ErrorValidationWrongMethod = errors.New("error validation scheme : not allowed method : WRONG")

	ConfigWithoutValidation = &generator.Config{
		ValidatePath:   false,
		ValidateQuery:  false,
		ValidateHeader: false,
		ValidateScheme: false,
	}
	ConfigWithValidation = &generator.Config{
		ValidatePath:   true,
		ValidateQuery:  true,
		ValidateHeader: true,
		ValidateScheme: true,
	}
)
