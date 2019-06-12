package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GrzegorzCzaprowski/go-exercises/rest-api/verifications"

	"github.com/GrzegorzCzaprowski/go-exercises/rest-api/models"
	"github.com/julienschmidt/httprouter"
)

func (h TodoHandler) Post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	todo := models.Todo{}
	err := json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		log.Println("error with decoding to json: ", err)
		return
	}
	err = verifications.CheckTodoContent(todo)
	if err != nil {
		log.Println("error with todo content: ", err)
		return
	}

	h.M.CreateTodo(todo)
}
