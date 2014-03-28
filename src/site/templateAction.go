package site;

import (
    "net/http"
    "../webserver"
);

// ---------------------- 
// Declarations

type TemplateAction struct {
    App     webserver.App
    page*   webserver.Page
}

type TemplateParams struct {
    Content string
}

func (a *TemplateAction) Init() {
    page, err := webserver.CreateDynamicPage("templatetest", &a.App.Router, "templates/test.htm");
    if (err != nil) {
        a.page = page;
    } else {
        panic(err);
    }
}

func (a *TemplateAction) Execute(writer http.ResponseWriter, request *http.Request, parameters map[string]string) (int, string) {
    result,_ := a.page.Render(TemplateParams{"Content to insert in the template."});
    return 200, result;
}

