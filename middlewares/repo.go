package middlewares

import (
	"fmt"
	"github.com/arickho/test_api/domain"
)

var currentID int
var Todos domain.Todos

// Give us some seed data
func Init() {
	RepoCreateTodo(domain.Todo{Name: "Write presentation"})
	RepoCreateTodo(domain.Todo{Name: "Host meetup"})
}

// RepoFindTodo -
func RepoFindTodo(ID int) domain.Todo {
	for _, t := range Todos {
		if t.ID == ID {
			return t
		}
	}

	// return empty Todo if not found
	return domain.Todo{}
}

// RepoCreateTodo -
func RepoCreateTodo(t domain.Todo) domain.Todo {
	currentID = currentID + 1
	t.ID = currentID
	Todos = append(Todos, t)
	return t
}

// RepoDestoryTodo -
func RepoDestoryTodo(ID int) error {
	for i, t := range Todos {
		if t.ID == ID {
			Todos = append(Todos[:i], Todos[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Could not find Todo with id of %d to delete", ID)
}
