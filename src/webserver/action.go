package webserver;

import (
    "net/http"
);

// ---------------------- 
// Declarations

type Action interface {
    Init();
    Execute(writer http.ResponseWriter, request *http.Request, parameters map[string]string) (int, string)
}
