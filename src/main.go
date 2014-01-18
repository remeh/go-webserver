package main;

import (
    "fmt"
    "./core"
    "./site"
);

// ---------------------- 

func main() {
    // instanciate the app
    var app core.App;

    app.Init();

    // example route
    app.Router.Add("index", &site.IndexAction{}, "/hello/:name");

    fmt.Println("[info] Starting the application.");
    app.Start(8080);
}

