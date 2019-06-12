package handlers

import "github.com/GrzegorzCzaprowski/go-exercises/rest-api/models"

type modelerUsers interface {
	CreateUser(user models.User) error
	LogUser(user models.User) error
}

type UserHandler struct {
	M modelerUsers
}
