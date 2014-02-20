package webserver;

import (
    "net/http"
);

// ---------------------- 
// Declarations

type Action interface {
    Init();
    Execute(request *http.Request, parameters map[string]string) string
}
