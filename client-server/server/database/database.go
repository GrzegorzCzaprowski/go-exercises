package database

import (
	"encoding/json"
	"io"
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
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}
	db.m.Lock()
	defer db.m.Unlock()
	db.Users[user.ID] = user
	err = json.NewEncoder(w).Encode(db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	user1, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}
	w.Write(user1)

}

func (db *Database) Delete(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	id, err := strconv.ParseUint(params["id"], 0, 64)
	if err != nil {
		log.Printf("%v\n", err)
	}
	db.m.Lock()
	delete(db.Users, id)
	err = json.NewEncoder(w).Encode(db.Users)
	if err != nil {
		log.Printf("%v\n", err)
	}
	db.m.Unlock()

}

func (db *Database) Get(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.ParseUint(params["id"], 0, 64)
	if err != nil {
		log.Printf("%v\n", err)
	}
	db.m.Lock()
	err = json.NewEncoder(w).Encode(db.Users[id])
	if err != nil {
		log.Printf("%v\n", err)
	}
	db.m.Unlock()

}

func (db *Database) SaveToFile(number int) {
	for {
		time.Sleep(time.Duration(number) * time.Second)
		file, err := json.MarshalIndent(db, "", " ")
		if err != nil {
			log.Printf("%v\n", err)
		}
		err = ioutil.WriteFile("database.json", file, 0644)
		if err != nil {
			log.Printf("%v\n", err)
		}
	}
}
