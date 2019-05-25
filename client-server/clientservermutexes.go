package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

type User struct {
	ID      uint64
	Name    string
	Surname string
	Email   string
}

type Database struct {
	m     sync.RWMutex
	Users map[uint64]User
}

func (db Database) Set(w http.ResponseWriter, req *http.Request) {
	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)
	db.m.Lock()
	db.Users[user.ID] = user
	json.NewEncoder(w).Encode(db)
	db.m.Unlock()
}

func (db Database) Delete(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.ParseUint(params["id"], 0, 64)
	db.m.Lock()
	db.Users[id] = User{}
	//rozmiar mapy nie zmieni sie, skasowany user zostanie zastąpiony pustym userem
	json.NewEncoder(w).Encode(db.Users)
	db.m.Unlock()
}

func (db Database) Get(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.ParseUint(params["id"], 0, 64)
	db.m.Lock()
	json.NewEncoder(w).Encode(db.Users[id])
	db.m.Unlock()
}

func main() {
	db := new(Database)
	db.Users = make(map[uint64]User)
	router := mux.NewRouter()

	router.HandleFunc("/users/", db.Set).Methods("POST")
	router.HandleFunc("/users/{id}/", db.Delete).Methods("DELETE")
	router.HandleFunc("/users/{id}/", db.Get).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
