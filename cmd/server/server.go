package main

import (
	"fmt"
	"golang-learning/config"
	"golang-learning/internal/foundation/configurator"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	config := config.NewConfig()
	configurator.ReadDefaultValues(config)
	configurator.ReadEnvironment(config)

	http.ListenAndServe(fmt.Sprintf("%s:%s", config.Application.Host, config.Application.Port), nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
