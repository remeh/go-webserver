package site

import (
	"fmt"
	"net/http"
)

// ----------------------
// Declarations

type IndexAction struct {
}

func (a *IndexAction) Init() {
}

func (a *IndexAction) Execute(writer http.ResponseWriter, request *http.Request, parameters map[string]string) (int, string) {
	return 200, fmt.Sprintf("hello " + parameters["name"])
}
