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
        panic(err);
    } else {
        a.page = page;
    }
}

func (a *TemplateAction) Execute(writer http.ResponseWriter, request *http.Request, parameters map[string]string) (int, string) {
    name := "No name given.";
    if (parameters["name"] != "") {
        name = parameters["name"];
    }
    result,_ := a.page.RenderNamedTemplate("test.htm", TemplateParams{name});
    return 200, result;
}
