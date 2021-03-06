package webserver

import (
	"fmt"
	"io/ioutil"
	"text/template"
)

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
	Type     string
	Template *template.Template
}

// ----------------------
// Methods

func CreateStaticPage(name string, path string) *Page {
	page := new(Page)
	page.Name = name
	page.Type = "STATIC"
	page.Template = nil

	content, err := ioutil.ReadFile(fmt.Sprintf("%s", path))
	if err != nil {
		fmt.Printf("x Error while creating the static page '%s' reading the file '%s'\n", page.Name, path)
	}

	page.Body = string(content)
	return page
}

func CreateDynamicPage(name string, router *Router, filenames ...string) (*Page, error) {
	page := new(Page)
	page.Name = name
	page.Type = "GOTEMPLATE"
	page.Template = nil

	funcMap := template.FuncMap{"rreverse": router.Reverse, "createmap": WebserverCreateMap}

	template, err := template.New(page.Name).Funcs(funcMap).ParseFiles(filenames...)

	if err != nil {
		fmt.Println("[error] Error while compiling template '%s' : %s", page.Name, err)
	} else {
		page.Template = template
	}

	return page, err
}

/**
 * Page initialization.
 */
func (p *Page) Init() {
	// Go Templates need to be compiled.
	if p.Type == "GOTEMPLATE" {
		template, err := template.New(p.Name).Parse(p.Body)
		if err != nil {
			fmt.Printf("Error while compiling template '%s' :\n %s\n", p.Name, err)
		} else {
			p.Template = template
		}
	}
	fmt.Printf(" - Page '%s' done\n", p.Name)
}

func (p *Page) Render(data interface{}) (string, error) {
	if p.Type == "GOTEMPLATE" {
		return p.RenderTemplate(data)
	}
	return p.Body, nil
}

func (p *Page) RenderNamedTemplate(templateName string, data interface{}) (string, error) {
	out := new(SimpleStringWriter)
	err := p.Template.ExecuteTemplate(out, templateName, data)
	if err != nil {
		fmt.Printf("Error while rendering the named template '%s' :\n %s\n", p.Name, err)
		return "", err
	}
	return out.Value, nil
}

func (p *Page) RenderTemplate(data interface{}) (string, error) {
	out := new(SimpleStringWriter)
	err := p.Template.Execute(out, data)
	if err != nil {
		fmt.Printf("Error while rendering the template '%s' :\n %s\n", p.Name, err)
		return "", err
	}
	return out.Value, nil
}
