package core;

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
);

// ---------------------- 
// Declarations

/**
 * An asset, ready each time needed to be rendered.
 * @author Rémy MATHIEU
 */
type Asset struct {
    // route of the asset (from the http request)
    Route string
    // how to read it on the disk
    Filename string
    // the data read
    Data []byte
}

// ----------------------
// Methods

// Looks into the asset directory with the given path
// for an asset. Returns nil if nothing is found.
func FindAsset(path string) *Asset {
    filename    := TransformAssetPath(path);
    data, err   := ioutil.ReadFile(filename);

    // No asset found.
    if (err != nil) {
        return nil;
    }

    // Creates and returns the asset
    var asset *Asset    = new(Asset);
    asset.Route         = path;
    asset.Filename      = filename;
    asset.Data          = data;

    return asset;
}

// ---------------------- 
// Class Methods

// Transforms the path of the request to an asset path.
// @param path      the path to transform.
func TransformAssetPath(path string) string {
    return fmt.Sprintf("public%s", path);
}

func (a *Asset) Render(w http.ResponseWriter, request *http.Request) {
    // set the type of the asset to the response header
    a.setResponseContentType(w.Header());

    // render the response
    fmt.Fprintf(w, "%s", a.Data);

    // log
    logAccess(request, false, "");
}

func (a *Asset) setResponseContentType(Header http.Header) {
    // suffix of the asset
    suffix := strings.ToLower(a.Filename[len(a.Filename)-3:]);

    // TODO configure an array of types in app/
    if (suffix == "css") {
        Header.Set("Content-type", "text/css");
    } else if (suffix == "json") {
        Header.Set("Content-type", "application/json");
    }
}

