package site;

import (
    "net/http"
    "../core"
);

// ---------------------- 
// Declarations

type TemplateAction struct {
    page* core.Page
}

func (a *TemplateAction) Init() {
    a.page = core.CreateDynamicPage("templatetest", "templates/test.htm");
}

func (a *TemplateAction) Execute(request *http.Request, parameters map[string]string) string {
    return a.page.Render("parameter for template");
}

