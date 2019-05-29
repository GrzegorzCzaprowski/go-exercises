package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"zadanka/go-exercises/client-server/server/database"

	"github.com/gorilla/mux"
)

func main() {
	var load bool
	fmt.Println("do you want to load database? y/n")

	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadByte()
		if input == 'y' {
			load = true
			break
		} else if input == 'n' {
			load = false
			break
		} else {
			println("wrong character, try again")
		}
	}

	db := new(database.Database)
	db.Users = make(map[uint64]database.User)

	if load {
		data, _ := ioutil.ReadFile("database.json")
		db = &database.Database{}
		_ = json.Unmarshal([]byte(data), &db)
	}
	router := mux.NewRouter()

	router.HandleFunc("/users/", db.Set).Methods("POST")
	router.HandleFunc("/users/{id}/", db.Delete).Methods("DELETE")
	router.HandleFunc("/users/{id}/", db.Get).Methods("GET")

	var flagServerAddress string
	flag.StringVar(&flagServerAddress, "addr", ":8000", "server address")
	var flagSavingDatabaseInterval int
	flag.IntVar(&flagSavingDatabaseInterval, "save", 5, "time beetween saves in seconds")
	flag.Parse()

	go db.SaveToFile(flagSavingDatabaseInterval)

	log.Fatal(http.ListenAndServe(flagServerAddress, router))
}
