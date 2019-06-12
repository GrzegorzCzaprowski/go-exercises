package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GrzegorzCzaprowski/go-exercises/rest-api/models"
	"github.com/julienschmidt/httprouter"
)

// func hashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	return string(bytes), err
// }

func (h UserHandler) PostUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	user := models.User{}
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		log.Println("error with decoding user to json: ", err)
		return
	}
	// var password string
	// password = user.Password
	// user.Password, err = hashPassword(password)
	// if err != nil {
	// 	log.Println("error with hashing user password: ", err)
	// 	return
	// }

	err = h.M.CreateUser(user)
	if err != nil {
		log.Println("error with creating user: ", err)
		return
	}
}
