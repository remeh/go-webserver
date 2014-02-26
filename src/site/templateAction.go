package site;

import (
    "net/http"
    "../webserver"
);

// ---------------------- 
// Declarations

type TemplateAction struct {
    page* webserver.Page
}

type TemplateParams struct {
    Content string
}

func (a *TemplateAction) Init() {
    a.page = webserver.CreateDynamicPage("templatetest", "templates/test.htm");
}

func (a *TemplateAction) Execute(writer http.ResponseWriter, request *http.Request, parameters map[string]string) (int, string) {
    return 200, a.page.Render(TemplateParams{"Content to insert in the template."});
}

