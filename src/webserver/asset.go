package webserver;

import (
    "fmt"
    "net/http"
    "io/ioutil"
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
    a.setResponseContentType(&w);

    // render the response
    fmt.Fprintf(w, "%s", a.Data);

    // log
    logAccess(request, false, "");
}

func (a *Asset) setResponseContentType(w *http.ResponseWriter) {
    writer := *w;
    writer.Header().Set("Content-type", http.DetectContentType(a.Data));

    // TODO checks that the Go http.DetectContentType is enough.
}

