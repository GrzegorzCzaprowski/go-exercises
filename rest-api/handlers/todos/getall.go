package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h TodoHandler) GetAll(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	todos, err := h.M.ReadAllTodos()
	if err != nil {
		log.Println("error with reading todos: ", err)
		w.WriteHeader(500)
		return
	}

	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		log.Println("error with encoding to json: ", err)
		w.WriteHeader(500)
		return
	}
}
