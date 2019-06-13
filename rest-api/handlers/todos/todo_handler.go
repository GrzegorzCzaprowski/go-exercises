package handlers

import "github.com/GrzegorzCzaprowski/go-exercises/rest-api/models"

type modelerTodos interface {
	CreateTodo(todo models.Todo) (models.Todo, error)
	ReadAllTodos() ([]models.Todo, error)
	ReadById(int) (models.Todo, error)
	UpdateById(todo models.Todo, id int) (models.Todo, error)
	RemoveById(id int) error
}
type TodoHandler struct {
	M modelerTodos
}
