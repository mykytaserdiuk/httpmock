package writer

import (
	"fmt"
	"net/http"
)

type Responser struct {
	Name string
}

func NewResponser(name string) *Responser {
	return &Responser{Name: name}
}
func (r *Responser) Error(w http.ResponseWriter, err error, status int) {
	message := fmt.Sprintf("Name: %s, error: %s, code: %d", r.Name, err.Error(), status)
	http.Error(w, message, status)
}
