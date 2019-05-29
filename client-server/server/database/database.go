package database

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

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

func (db *Database) Set(w http.ResponseWriter, req *http.Request) {
	var user User
	json.NewDecoder(req.Body).Decode(&user)
	db.m.Lock()
	db.Users[user.ID] = user
	json.NewEncoder(w).Encode(db)
	db.m.Unlock()
}

func (db *Database) Delete(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.ParseUint(params["id"], 0, 64)
	if err != nil {
		log.Println("Error: can't parse string to integer!")
	}
	db.m.Lock()
	delete(db.Users, id)
	json.NewEncoder(w).Encode(db.Users)
	db.m.Unlock()
}

func (db *Database) Get(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.ParseUint(params["id"], 0, 64)
	if err != nil {
		log.Println("Error: can't parse string to integer!")
	}
	db.m.Lock()
	json.NewEncoder(w).Encode(db.Users[id])
	db.m.Unlock()
}

func (db *Database) SaveToFile(number int) {
	for {
		time.Sleep(time.Duration(number) * time.Second)
		file, err := json.MarshalIndent(db, "", " ")
		if err != nil {
			log.Println("Error: can't save database to file!")
		}
		ioutil.WriteFile("database.json", file, 0644)
	}
}
