remeh's go webserver
==

Basic Go webserver used for the [remy.io](http://remy.io) website.

Usage:
  * 1st method: build an executable :`go build src/main.go && ./main`
  * 2nd method: `go run src/main.go`

Features :
  * **Page routing configuration** : ['/','/index'] -> index.htm, etc.
  * **Assets support** : the directory assets is directly rendered without evaluation

Roadmap :
  * **Full list of asset mime-types**
  * **Support go template** to render pages
