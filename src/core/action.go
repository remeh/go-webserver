package core;

import (
    "net/http"
);

// ---------------------- 
// Declarations

type Action interface {
    Execute(request *http.Request, parameters map[string]string) string
}
