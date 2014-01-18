package core;

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "text/template"
);

// ---------------------- 
// Declarations

/**
 * A page of the app.
 * Stored in RAM.
 *
 * @author Rémy MATHIEU
 */
type Page struct {
    Name string
    Body string
    // Possible values :
    // STATIC, GOTEMPLATE
    Type string
    Template *template.Template
}

// ---------------------- 
// Methods

func CreateStaticPage(name string, path string) *Page {
    page := new(Page);
    page.Name       = name;
    page.Type       = "STATIC";
    page.Template   = nil;

    content, err := ioutil.ReadFile(fmt.Sprintf("%s", path));
    if (err != nil) {
        fmt.Printf("x Error while creating the static page '%s' reading the file '%s'\n", page.Name, path);
    }

    page.Body       = string(content);
    return page;
}

/**
 * Page initialization.
 */
func (p *Page) Init() {
    // Go Templates need to be compiled.
    if (p.Type == "GOTEMPLATE") {
        template, err := template.New(p.Name).Parse(p.Body);
        if (err != nil) {
            fmt.Println("Error while compiling template '%s' : %s", p.Name, err);
        } else {
            p.Template = template;
        }
    }
    fmt.Printf(" - Page '%s' done\n",p.Name);
}

func (p *Page) Render(w http.ResponseWriter, request *http.Request) {
    if (p.Type == "GOTEMPLATE") {
        p.renderGoTemplate(w, request);
    } else {
        fmt.Fprintf(w, "%s", p.Body);
    }
}

func (p *Page) renderGoTemplate(w http.ResponseWriter, request *http.Request, params... string) {
    err := p.Template.Execute(w, params);
    if (err != nil) {
        fmt.Println("Error while rendering the template '%s' : %s", p.Name, err);
    }
}
