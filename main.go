package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arickho/test_api/middlewares"
	"github.com/arickho/test_api/server"
)

func main() {

	fmt.Printf("Initializing api on localhost:8080")

	middlewares.Init()
	router := server.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
