package handlers

import "github.com/GrzegorzCzaprowski/go-exercises/rest-api/models"

type modelerTodos interface {
	CreateTodo(todo models.Todo) error
	ReadAllTodos() ([]models.Todo, error)
	ReadById(int) (models.Todo, error)
	UpdateById(todo models.Todo, id int) error
	RemoveById(id int) error
}
type TodoHandler struct {
	M modelerTodos
}