package main

import (
	"fmt"
	"go-rest-api/router"
	"log"
	"net/http"
)

func main() {
	const port = 8080
	router := http.HandlerFunc(router.Serve)
	fmt.Printf("listening on port %d \n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
