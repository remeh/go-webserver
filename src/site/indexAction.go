package site;

import (
    "fmt"
    "net/http"
);

// ---------------------- 
// Declarations

type IndexAction struct {
}

func (a *IndexAction) Init() {
}

func (a *IndexAction) Execute(request *http.Request, parameters map[string]string) string {
    return fmt.Sprintf("hello " + parameters["name"]);
}
