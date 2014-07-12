package main

import (
	"fmt"
	"github.com/remeh/go-webserver"

	"./site"
)

// ----------------------

func main() {
	// instanciate the app
	var app webserver.App

	app.Init()

	// example route
	app.Router.Add("index", "*", &site.IndexAction{}, "/hello/:name")
	app.Router.Add("template", "POST", &site.TemplateAction{App: app}, "/template/:name")

	fmt.Println("[info] Starting the application.")
	app.Start(8080)
}
