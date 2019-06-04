package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

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
