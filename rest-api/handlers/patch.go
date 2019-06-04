package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/GrzegorzCzaprowski/go-exercises/rest-api/models"
	"github.com/julienschmidt/httprouter"
)

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
