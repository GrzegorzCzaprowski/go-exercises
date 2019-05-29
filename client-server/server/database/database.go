package database

import (
	"encoding/json"
	"io/ioutil"
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
	_ = json.NewDecoder(req.Body).Decode(&user)
	db.m.Lock()
	db.Users[user.ID] = user
	json.NewEncoder(w).Encode(db)
	db.m.Unlock()
}

func (db *Database) Delete(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.ParseUint(params["id"], 0, 64)
	db.m.Lock()
	delete(db.Users, id)
	json.NewEncoder(w).Encode(db.Users)
	db.m.Unlock()
}

func (db *Database) Get(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.ParseUint(params["id"], 0, 64)
	db.m.Lock()
	json.NewEncoder(w).Encode(db.Users[id])
	db.m.Unlock()
}

func (db *Database) SaveToFile(number int) {
	for {
		time.Sleep(time.Duration(number) * time.Second)
		file, _ := json.MarshalIndent(db, "", " ")
		_ = ioutil.WriteFile("database.json", file, 0644)
	}
}
