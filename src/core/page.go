package core;

import (
    "fmt"
    "net/http"
);

// ---------------------- 
// Declarations

type PageInterface interface {
    Render(w http.ResponseWriter, request *http.Request, route string) string
}

/**
 * A page of the app.
 * Stored in RAM.
 * @author Rémy MATHIEU
 */
type Page struct {
    Name string
    Body string
}

// ---------------------- 
// Methods

/**
 * Page initialization.
 */
func (p *Page) Init() {
    fmt.Printf(" - Page '%s' OK\n",p.Name);
}

func (p *Page) Render(w http.ResponseWriter, request *http.Request){
    fmt.Fprintf(w, "%s", p.Body);
    logAccess(request, false, "");
}
