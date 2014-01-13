package core;

import (
    "fmt"
    "io/ioutil"
    "net/http"
);

// ---------------------- 
// Declarations

/**
 * The router struct.
 * @author Rémy MATHIEU
 */
type Router struct {
    // dynamic pages
    Pages map[string]*Page;
}

/**
 * Configuration format.
 * TODO maybe move this in a different go file.
 */
type ConfigurationFormat struct {
    // The different pages
    Pages   []PageConfigurationFormat
}

/**
 * Page configuration format.
 * TODO maybe move this in a different go file.
 */
type PageConfigurationFormat struct {
    Routes  []string
    Name    string
    File    string
    // Page type, possible values STATIC / GOTEMPLATE
    Type    string
}

// ---------------------- 
// Public methods

/**
 * Router initialization.
 */
func (r *Router) Init() {
    r.Pages = make(map[string]*Page);
    fmt.Println(" - Router init OK");
}


func (r *Router) Start() {
    http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) { r.route(w, request) });
    fmt.Println(" - Router started");
}

// ---------------------- 
// Private methods

/**
 * Routing between pages and assets. Or 404.
 */
func (r *Router) route(w http.ResponseWriter, request *http.Request) {
    // Is this an existing page ?
    page := r.Pages[request.URL.Path];

    // Nope
    if (page == nil) {
        asset := FindAsset(request.URL.Path);
        // assets ?
        if (asset != nil) {
            asset.Render(w, request);
            return;
        }

        // Nope.
        // TODO 404 page
        w.WriteHeader(404);
        fmt.Fprint(w, "404");
        logAccess(request, true, "-> 404");
        return;
    }

    // render the page
    page.Render(w, request);
}

/**
 * Evaluates the router part of the configuration.
 * @param config        the read configuration
 */
func (r *Router) evalutateConfiguration(config ConfigurationFormat) {
    for i := 0; i < len(config.Pages); i++ {
        p := config.Pages[i];

        // Read the content of the page / template.

        content, err := ioutil.ReadFile(fmt.Sprintf("pages/%s", p.File)); // TODO isn't there a security issue there?
        if (err != nil) {
            fmt.Printf(" x Error while loading the page '%s' in the file '%s' : %s \n", p.Name, p.File, err);
            continue;
        }

        // Creates the page.

        var page *Page  = new(Page);
        page.Body       = string(content);
        page.Name       = p.Name;
        page.Type       = p.Type;

        // Prepares the page

        page.Init();

        // Adds it to the router.
        // TODO regexp etc.

        for j := 0; j < len(p.Routes); j++ {
            r.Pages[p.Routes[j]] = page;
        }

        fmt.Printf(" - %s Page '%s' loaded (with: '%s', binded on %s)\n", p.Type, p.Name, p.File, p.Routes);
    }
}
