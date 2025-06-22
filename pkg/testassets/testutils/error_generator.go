package testutils

import "fmt"

const (
	errResponseValidationNotAllowedCode = `error validation scheme : response is not valid : not allowed response status`
)

func ErrResponseValidationNotAllowedCode(code int) error {
	return fmt.Errorf("%s : %d", errResponseValidationNotAllowedCode, code)
}
