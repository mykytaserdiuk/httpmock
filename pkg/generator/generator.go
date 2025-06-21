package generator

import "github.com/gorilla/mux"

type Generator struct {
	router *mux.Router
}

func NewGenerator() *Generator {
	return &Generator{router: mux.NewRouter()}
}
