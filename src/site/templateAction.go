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

func (a *TemplateAction) Execute(request *http.Request, parameters map[string]string) string {
    return a.page.Render(TemplateParams{"Content to insert in the template."});
}

