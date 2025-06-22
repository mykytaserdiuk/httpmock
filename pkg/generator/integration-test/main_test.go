package integration_test

import (
	"log"
	"testing"
)

func TestMain(t *testing.M) {
	code := t.Run()
	log.Print(code)
	initAllExample()
}

func initAllExample() {
	initGoodExample()
}

func initGoodExample() {
	//runner := generator.NewRunner(ta.ConfigWithValidation)
	//path := "./example/good-example.yaml"
	//var mock models.MockScheme
	//if err := cfg.UnmarshalYAML(path, &mock); err != nil {
	//	log.Fatal(err)
	//}
	//if err := runner.Launch(&mock); err != nil {
	//	log.Fatal(err)
	//}
}
