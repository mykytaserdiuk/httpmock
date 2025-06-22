package main

import (
	"flag"
	"fmt"

	"github.com/mykytaserdiuk9/httpmock/pkg/cfg"
	"github.com/mykytaserdiuk9/httpmock/pkg/generator"
	"github.com/mykytaserdiuk9/httpmock/pkg/models"
)

var (
	noPathValidation   = flag.Bool("no-path", false, "Disable path validation")
	noQueryValidation  = flag.Bool("no-query", false, "Disable query validation")
	noHeaderValidation = flag.Bool("no-header", false, "Disable header validation")
	noSchemeValidation = flag.Bool("no-scheme", false, "Disable scheme validation")
)

func main() {
	flag.Parse()

	var mock models.MockScheme
	config := &generator.Config{
		ValidatePath:   !*noPathValidation,
		ValidateQuery:  !*noQueryValidation,
		ValidateHeader: !*noHeaderValidation,
		ValidateScheme: !*noSchemeValidation,
	}
	err := cfg.UnmarshalYAML("./cmd/example.yaml", &mock)

	runner := generator.NewRunner(config)
	err = runner.Launch(&mock)
	fmt.Print(err)
}
