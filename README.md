remeh's go webserver
==

Powerful but simple Go webserver.

Benchmark using [wrk](https://github.com/wg/wrk), it reaches more than 41000req/s on a dynamic route (with a parameter directly in the URL) rendering a simple text.

## Run the example
  * 1st method: build an executable :`go build example/main.go && ./main`
  * 2nd method: `go run example/main.go`

## Features
  * **Routing with regexp** : ex: GET ['/list/:id','/l/:id'] can send to an action GET which receives a parameter 'id',
  * **Support go template** : to render pages in Go templates,
  * **Reverse routing** : allows to rewrite routes directly in a template,
  * **Assets support** : the directory assets is directly rendered without evaluation.

## Roadmap
  * **Use regexp only when needed** : route without parameters could be faster if no regexp were used for them.

## How-to

### Get the package

To get the `webserver` package :

```
go get github.com/remeh/go-webserver
```

You're now able to :

```
import "github.com/remeh/go-webserver"
```

in your source code.

See `examples/main.go` and the content of `examples/site` for an example of usage.
