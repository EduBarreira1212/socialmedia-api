package main

import (
	"fmt"
	"net/http"
	"socialmedia-api/src/config"
	"socialmedia-api/src/router"
)

func main() {
	config.Load()
	fmt.Printf("Running server on port %d!", config.Port)
	r := router.Generate()

	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r)
}
