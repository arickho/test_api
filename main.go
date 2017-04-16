package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
	"test_api/models"
)

func main() {
	fmt.Printf("Initializing api on localhost:8080")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// Index - api splash
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// TodoIndex - the todo splash
func TodoIndex(w http.ResponseWriter, r *http.Request) {

	todos := todo.Todos{
		todo.Todo{Name: "Write presentation"},
		todo.Todo{Name: "Host meetup "},
	}

	json.NewEncoder(w).Encode(todos)
}

// TodoShow - todo detail
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
