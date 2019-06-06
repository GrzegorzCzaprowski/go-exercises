package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GrzegorzCzaprowski/go-exercises/rest-api/verifications"

	"github.com/GrzegorzCzaprowski/go-exercises/rest-api/models"
	"github.com/julienschmidt/httprouter"
)

func (h Handler) Post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	todo := models.Todo{}
	err := json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		log.Panicln("error with decoding to json: ", err)
	}
	err = verifications.CheckTodoContent(todo)
	if err != nil {
		log.Panicln("error with todo content: ", err)
	}

	h.M.CreateTodo(todo)
}
