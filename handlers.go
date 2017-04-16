package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"test_api/models"
)

// Index - api splash
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

// TodoIndex - the todo splash
func TodoIndex(w http.ResponseWriter, r *http.Request) {

	todos := todo.Todos{
		todo.Todo{Name: "Write presentation"},
		todo.Todo{Name: "Host meetup "},
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

// TodoShow - todo detail
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoID)
}
