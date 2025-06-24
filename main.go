package main

import (
	"fmt"
	"net/http"
	"socialmedia-api/src/router"
)

func main() {
	fmt.Println("Running server on port 5000!")
	r := router.Generate()

	http.ListenAndServe(":5000", r)
}
