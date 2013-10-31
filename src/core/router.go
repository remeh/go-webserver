package core;

import (
    "fmt"
    "encoding/json"
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
}

// ---------------------- 
// Public methods

/**
 * Router initialization.
 */
func (r *Router) Init() {
    r.Pages = make(map[string]*Page);
    r.readConfiguration();
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
 * Reads the router configuration.
 */
func (r *Router) readConfiguration() {
    /*
     * Read the file
     */

    data, err := ioutil.ReadFile("app/routes.json"); // XXX hardcoded configuration
    if (err != nil) {
        fmt.Printf(" x Unable to read the router configuration : error while reading the file : \n%s\n",err);
        return;
    }

    /*
     * Unmarshal the json.
     */

    var config ConfigurationFormat;
    err = json.Unmarshal(data, &config);
    if (err != nil) {
        fmt.Printf(" x Unable to read the router configuration : error while unmarshaling the data : \n%s\n",err);
    }

    /*
     * Evaluate the configuration.
     */

    r.evalutateConfiguration(config);
}


/**
 * Evaluates the read configuration.
 * @param config        the read configuration
 */
func (r *Router) evalutateConfiguration(config ConfigurationFormat) {
    for i := 0; i < len(config.Pages); i++ {
        p := config.Pages[i];

        // TODO move this in page.go

        /*
         * Read the content of the template.
         */

        content, err := ioutil.ReadFile(fmt.Sprintf("pages/%s", p.File)); // TODO isn't there a security issue there?
        if (err != nil) {
            fmt.Printf(" x Error while loading the page '%s' in the file '%s'\n", p.Name, p.File);
            continue;
        }

        /*
         * Creates the page.
         */

        var page *Page  = new(Page);
        page.Body       = string(content);
        page.Name       = p.Name;

        /*
         * Adds it to the router.
         */

        for j := 0; j < len(p.Routes); j++ {
            r.Pages[p.Routes[j]] = page;
        }

        fmt.Printf(" - Page '%s' loaded (with: '%s', binded on %s)\n", p.Name, p.File, p.Routes);
    }
}
