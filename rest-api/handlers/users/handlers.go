package handlers

import (
	"github.com/GrzegorzCzaprowski/go-exercises/rest-api/models"
)

type modeler interface {
	CreateUser(user models.User) error
	LogUser(user models.User) error
}

type Handler struct {
	M modeler
}
