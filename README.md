remeh's go webserver
==

Basic Go webserver used for the [remy.io](http://remy.io) website.

Usage:
  * 1st method: build an executable :`go build src/main.go && ./main`
  * 2nd method: `go run src/main.go`

Features :
  * **Routing with regexp** : ex: GET ['/list/:id','/l/:id'] can send to an action GET which receives a parameter 'id',
  * **Support go template** : to render pages in Go templates,
  * **Reverse routing** : allows to rewrite routes directly in a template,
  * **Assets support** : the directory assets is directly rendered without evaluation.

Roadmap :
  * **Use regexp only when needed** : route without parameters could be faster if no regexp were use for them.
