package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Initializing api on localhost:8080")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
