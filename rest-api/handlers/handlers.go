package handlers

import (
	"github.com/GrzegorzCzaprowski/go-exercises/rest-api/models"
)

type modeler interface {
	CreateTodo(todo models.Todo) error
	ReadAllTodos() ([]models.Todo, error)
	ReadById(int) (models.Todo, error)
	UpdateById(todo models.Todo, id int) error
	RemoveById(id int) error
	CreateUser(user models.User) error
	LogUser(user models.User) error
}

type Handler struct {
	M modeler
}
