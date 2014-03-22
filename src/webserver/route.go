package webserver;

import (
    "fmt"
    "regexp"
    "strings"
);

// ----------------------
// Declarations

/**
 * A route configured for the webapp.
 *
 * @author Rémy MATHIEU
 */
type Route struct {
    route   string;
    method  string;
    expr    *regexp.Regexp;
    params  map[int]string;
}

// ----------------------
// Methods

/**
 * Page initialization.
 */
func (r *Route) Init(method string, route string) {
    r.params = make(map[int]string);

    // Looks for parameters in the route
    varsRegexp := regexp.MustCompile(":[a-zA-Z]*");
    finalRoute := route;

    // Remembers their position in the route and replace them with the regexp (.*)
    params := varsRegexp.FindAllString(route, -1);
    for i := range params {
        finalRoute = strings.Replace(finalRoute, params[i], "(.*)", 1);
        r.params[i] = strings.Trim(params[i],":");
    }

    // Compiles the regexp.
    expr, err     := regexp.Compile(fmt.Sprintf("^%s$",finalRoute));

    // Stores the information.
    r.expr      = expr;
    r.route     = route;
    r.method    = method;

    if (err != nil) {
        fmt.Printf("[error] Error while compiling the route %s :\n", route);
        fmt.Println(err);
    } else {
        fmt.Printf("[info] Route %s '%s' compiled.\n", method, route);
    }
}

/**
 * Returns whether the given route string match the current Route.
 * @param route     the route to match
 * @return true if the given route string matches.
 */
func (r *Route) Match(method string, route string) bool {
    return (r.method == "*" || strings.Contains(r.method, method)) && r.expr.MatchString(route);
}

/**
 * Extracts the parameters from the given route string.
 * @param route     the route string from which we want to extract the parameters.
 * @param the extracted parameters.
 */
func (r *Route) ReadParameters(route string) map[string]string {
    params := make(map[string]string);
    submatches := r.expr.FindStringSubmatch(route);
    for i := 1; i < len(submatches); i++ {
        params[ r.params[i-1] ] = submatches[i];
    }
    return params;
}
