package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (h Handler) Delete(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		log.Println("cant parse paramater to int: ", err)
		return
	}

	err = h.M.RemoveById(id)
	if err != nil {
		log.Println("error with updating todo: ", err)
		return
	}
}
