remeh's go webserver
==

Basic Go webserver used for the [remy.io](http://remy.io) website.

Usage:
  * 1st method: build an executable :`go build src/main.go && ./main`
  * 2nd method: `go run src/main.go`

Features :
  * **Routing with regexp** : ['/list/:id','/l/:id'] -> action list which receive a parameter 'id'
  * **Assets support** : the directory assets is directly rendered without evaluation
  * **Support go template** : to render pages

