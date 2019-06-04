package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GrzegorzCzaprowski/go-exercises/rest-api/models"
	"github.com/julienschmidt/httprouter"
)

type modeler interface {
	CreateTodo(todo models.Todo) error
	ReadAllTodos() ([]models.Todo, error)
}

type Handler struct {
	M modeler
}

func (h Handler) Post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	todo := models.Todo{}
	err := json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		log.Panicln("error with decoding to json: ", err)
	}

	h.M.CreateTodo(todo)
}

func (h Handler) GetAll(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	todos, err := h.M.ReadAllTodos()
	if err != nil {
		log.Panicln("error with reading todos: ", err)
	}
	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		log.Panicln("error with encoding to json: ", err)
	}
}
