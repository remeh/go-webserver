remeh's go webserver
==

Basic Go webserver used for the [remy.io](http://remy.io) website.

Usage:
  * 1st method: build an executable :`go build src/main.go && ./main`
  * 2nd method: `go run src/main.go`

Features :
  * **Routing with regexp** : ex: ['/list/:id','/l/:id'] can send to an action which receives a parameter 'id'
  * **Support go template** : to render pages
  * **Assets support** : the directory assets is directly rendered without evaluation

