package todo

import "time"

// Todo - Model
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos - List of todo's
type Todos []Todo
