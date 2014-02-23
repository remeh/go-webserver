package webserver;

import (
    "fmt"
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
    Actions map[*Route]Action;
}

// ---------------------- 
// Public methods

/**
 * Router initialization.
 */
func (r *Router) Init() {
    r.Actions = make(map[*Route]Action);
    fmt.Println("[info] Router init OK");
}

/**
 * Starts the router.
 */
func (r *Router) Start() {
    http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) { r.route(w, request) });
    fmt.Println("[info] Router started");
}

/**
 * Adds a named route to the router.
 * @param name      the name of the route
 * @param action    the action to execute.
 * @param routes    on which routes this route is rendered.
 * TODO post method ?
 */
func (r* Router) Add(name string, action Action, routes... string) {
    for j := 0; j < len(routes); j++ {
        route := routes[j];

        // Creates the route.
        newRoute := new(Route);
        newRoute.Init(route);

        // Inits the action
        action.Init();

        // stores the action
        r.Actions[newRoute] = action;
    }
}

// ---------------------- 
// Private methods

/**
 * Routing between pages and assets. Or 404.
 */
func (r *Router) route(w http.ResponseWriter, request *http.Request) {
    // Look for an existing route.
    url     := request.URL.Path;
    route   := r.matchRoute(url);

    if (route != nil) {
        // executes the action
        go logAccess(request, false, "");

        params := route.ReadParameters(url);
        action := r.Actions[route];

        httpCode, body := action.Execute(request, params);
        w.WriteHeader(httpCode);
        fmt.Fprintf(w, "%s", body);
        return;
    }

    // nope, looks in assets
    asset := FindAsset(request.URL.Path);
    // assets ?
    if (asset != nil) {
        asset.Render(w, request);
        return;
    }

    // nope.
    // TODO 404 page
    w.WriteHeader(404);
    fmt.Fprint(w, "404");
    go logAccess(request, true, "-> 404");
    return;
}

/**
 * Looks whether a route is matching the given pattern.
 * @param route     the pattern to match
 * @return the action to execute if some found.
 */
func (r *Router) matchRoute(url string) *Route {
    // Look through the whole route if one matches
    for key, _ := range r.Actions {
        if (key.Match(url)) {
            // This one match! Return the route.
            return key;
        }
    }
    return nil;
}