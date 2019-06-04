package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/GrzegorzCzaprowski/go-exercises/rest-api/models"
	"github.com/julienschmidt/httprouter"
)

type modeler interface {
	CreateTodo(todo models.Todo) error
	ReadAllTodos() ([]models.Todo, error)
	ReadById(int) (models.Todo, error)
	UpdateById(todo models.Todo, id int) error
	RemoveById(id int) error
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

func (h Handler) Get(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		log.Println("cant parse paramater to int: ", err)
	}
	todo, err := h.M.ReadById(id)
	if err != nil {
		log.Panicln("error with reading todo: ", err)
	}
	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		log.Panicln("error with encoding to json: ", err)
	}
}

func (h Handler) Patch(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		log.Println("cant parse paramater to int: ", err)
	}

	todo := models.Todo{}
	err = json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		log.Panicln("error with decoding to json: ", err)
	}

	err = h.M.UpdateById(todo, id)
	if err != nil {
		log.Panicln("error with updating todo: ", err)
	}
}

func (h Handler) Delete(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		log.Println("cant parse paramater to int: ", err)
	}

	err = h.M.RemoveById(id)
	if err != nil {
		log.Panicln("error with updating todo: ", err)
	}
}
