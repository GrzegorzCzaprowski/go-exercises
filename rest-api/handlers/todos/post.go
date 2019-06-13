package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GrzegorzCzaprowski/go-exercises/rest-api/models"
	"github.com/julienschmidt/httprouter"
)

func (h TodoHandler) Post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	todo := models.Todo{}
	err := json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		log.Println("error with decoding to json: ", err)
		w.WriteHeader(500)
		return
	}

	todo, err = h.M.CreateTodo(todo)
	if err != nil {
		log.Println("error with creating todo: ", err)
		w.WriteHeader(500)
		return
	}

	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		log.Println("error with encoding to json: ", err)
		w.WriteHeader(500)
		return
	}
}
