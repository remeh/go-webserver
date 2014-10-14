package webserver

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// ----------------------
// Declarations

const (
	// Days before expiration of the image
	DAYS_BEFORE_EXPIRATION = 30
	// Date format for headers of the response
	DATE_FORMAT = "Mon, 02 Jan 2006 15:04:05 MST"
)

//
// An asset, ready each time needed to be rendered.
// @author Rémy MATHIEU
//
type Asset struct {
	// route of the asset (from the http request)
	Route string
	// how to read it on the disk
	Filename string
	// the data read
	Data []byte
	// last modified-date
	LastModified string
}

// ----------------------
// Methods

// Looks into the asset directory with the given path
// for an asset. Returns nil if nothing is found.
func FindAsset(path string) *Asset {
	filename := TransformAssetPath(path)
	data, err := ioutil.ReadFile(filename)

	// No asset found.
	if err != nil {
		return nil
	}

	// Creates and returns the asset
	var asset *Asset = new(Asset)
	asset.Route = path
	asset.Filename = filename
	asset.Data = data

	// Sets the last modified date if readable.
	file, err := os.Open(filename)
	if err != nil {
		LogWebserverInfof("Unable to read information on the file '%s'.", filename)
		return nil
	}

	stats, err := file.Stat()
	if err == nil {
		asset.LastModified = stats.ModTime().Format(DATE_FORMAT)
	} else {
		asset.LastModified = time.Now().Format(DATE_FORMAT)
	}

	return asset
}

// ----------------------
// Class Methods

// Transforms the path of the request to an asset path.
// @param path      the path to transform.
func TransformAssetPath(path string) string {
	return fmt.Sprintf("public%s", path)
}

func (a *Asset) Render(w http.ResponseWriter, request *http.Request) {
	// set the type of the asset to the response header
	a.setResponseContentType(&w)

	// render the response
	fmt.Fprintf(w, "%s", a.Data)

	// log
	logAccess(request, false, "")
}

func (a *Asset) setResponseContentType(w *http.ResponseWriter) {
	writer := *w

	writer.Header().Set("Content-type", http.DetectContentType(a.Data))
	writer.Header().Set("Last-Modified", a.LastModified)

	// http.DetectContentType is not enough at all.
	suffix := strings.ToLower(a.Filename[len(a.Filename)-4:])

	// TODO configure an array of types in app/
	if suffix == ".css" {
		writer.Header().Set("Content-type", "text/css")
	} else if suffix == "json" {
		writer.Header().Set("Content-type", "application/json")
	}
}
