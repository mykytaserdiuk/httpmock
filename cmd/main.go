package main

import (
	"flag"
	"fmt"

	"github.com/mykytaserdiuk9/httpmock/pkg/generator"
	"github.com/mykytaserdiuk9/httpmock/pkg/schema"
)

var (
	noPathValidation   = flag.Bool("no-path", false, "Disable path validation")
	noQueryValidation  = flag.Bool("no-query", false, "Disable query validation")
	noHeaderValidation = flag.Bool("no-header", false, "Disable header validation")
	noSchemeValidation = flag.Bool("no-scheme", false, "Disable scheme validation")
	schemePath         = flag.String("path", "./api/example.yaml", "Path to scheme")
)

func main() {
	flag.Parse()
	config := &generator.Config{
		ValidatePath:   !*noPathValidation,
		ValidateQuery:  !*noQueryValidation,
		ValidateHeader: !*noHeaderValidation,
		ValidateScheme: !*noSchemeValidation,
	}

	mock, err := schema.Get(*schemePath)
	if err != nil {
		fmt.Println(err)
	}
	HTTPServer := generator.NewUpper(mock.Port)
	runner := generator.NewRunner(config, HTTPServer)
	if err := runner.Launch(mock); err != nil {
		fmt.Println(err)
	}
}
