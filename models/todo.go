package todo

import "time"

// Todo - Model
type Todo struct {
	Name      string
	Completed bool
	Due       time.Time
}

// Todos - List of todo's
type Todos []Todo
