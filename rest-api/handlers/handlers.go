package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GrzegorzCzaprowski/go-exercises/rest-api/models"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	S models.Server
}

func (h Handler) Post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	todo := &models.Todo{}
	err := json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		log.Println("error with decoding to json: ", err)
	}
	defer req.Body.Close()

	h.S.CreateTodo(todo)
}

func (h Handler) GetAll(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	todos, err := h.S.ReadAllTodos()
	if err != nil {
		log.Println("error with reading todos: ", err)
	}
	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		log.Println("error with encoding to json: ", err)
	}
}
