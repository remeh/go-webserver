remeh's go webserver
==

Basic Go webserver used for [blurmbl.com](http://blurmbl.com), the [remy.io](http://remy.io) website and some others.

## Usage
  * 1st method: build an executable :`go build src/main.go && ./main`
  * 2nd method: `go run src/main.go`

## Features
  * **Routing with regexp** : ex: GET ['/list/:id','/l/:id'] can send to an action GET which receives a parameter 'id',
  * **Support go template** : to render pages in Go templates,
  * **Reverse routing** : allows to rewrite routes directly in a template,
  * **Assets support** : the directory assets is directly rendered without evaluation.

## Roadmap
  * **Use regexp only when needed** : route without parameters could be faster if no regexp were use for them.
  * **Correct package for Go** : the webserver pkg shouldn't be inside src/webserver

## How-to

### Get the package

To get the `webserver` package :

```
go get github.com/remeh/go-webserver/src/webserver
```

You're now able to :

```
import github.com/remeh/go-webserver/src/webserver
```

in your source code.

See `src/main.go` and the content of `src/site` for an example of usage.
