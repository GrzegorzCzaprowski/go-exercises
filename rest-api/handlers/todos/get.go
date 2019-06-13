package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (h TodoHandler) Get(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		log.Println("cant parse paramater to int: ", err)
		w.WriteHeader(500)
		return
	}
	todo, err := h.M.ReadById(id)
	if err != nil {
		log.Println("error with reading todo: ", err)
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
