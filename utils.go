package webserver

import (
    "fmt"
	"log"
)

// That method is used to fulfill a big miss of Go templates : creation of maps in the templates.
// The number of parameters given to this method MUST BE pair :
// pair value is used as a key in the map, next  impaired value is used as value its the map.
func WebserverCreateMap(parameters ...string) map[string]string {
	m := make(map[string]string)
	if len(parameters)%2 != 0 {
		return m
	}
	for i := 0; i < len(parameters); i += 2 {
		m[parameters[i]] = parameters[i+1]
	}
	return m
}

func LogWebserverInfo(s string) {
	log.Println("[webserver] [info] " + s)
}

func LogWebserverInfof(s string, a ...interface{}) {
	log.Printf("[webserver] [info] %s", fmt.Sprintf(s, a...))
}

func LogWebserverErrorf(s string, a ...interface{}) {
	log.Printf("[webserver] [info] %s", fmt.Sprintf(s, a))
}

func LogWebserverError(e string) {
	log.Println("[webserver] [error] " + e)
}
