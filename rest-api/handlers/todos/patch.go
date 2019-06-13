package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/GrzegorzCzaprowski/go-exercises/rest-api/models"
	"github.com/julienschmidt/httprouter"
)

func (h TodoHandler) Patch(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		log.Println("cant parse paramater to int: ", err)
		w.WriteHeader(500)
		return
	}

	todo := models.Todo{}
	err = json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		log.Println("error with decoding to json: ", err)
		w.WriteHeader(500)
		return
	}

	todo, err = h.M.UpdateById(todo, id)
	if err != nil {
		log.Println("error with updating todo: ", err)
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
