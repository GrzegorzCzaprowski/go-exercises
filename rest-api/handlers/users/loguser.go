package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GrzegorzCzaprowski/go-exercises/rest-api/models"
	"github.com/julienschmidt/httprouter"
)

func (h UserHandler) LogUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	user := models.User{}
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		log.Println("error with decoding user to json: ", err)
		w.WriteHeader(500)
		return
	}
	user, err = h.M.LogUser(user)
	if err != nil {
		log.Println("error with logging user: ", err)
		w.WriteHeader(500)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println("error with encoding to json: ", err)
		w.WriteHeader(500)
		return
	}
}
