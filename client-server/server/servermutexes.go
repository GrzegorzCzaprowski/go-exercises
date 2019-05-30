package main

//errory obługujemy w handlerach gorutynach i mainie
import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/GrzegorzCzaprowski/go-exercises/client-server/server/database"
	"github.com/gorilla/mux"
)

func loadDatabase() bool {
	for {
		var input string
		fmt.Scan(&input)
		if input == "y" {
			return true
		} else if input == "n" {
			return false
		} else {
			fmt.Println("wrong character, try again")
		}
	}
}

func main() {
	fmt.Println("do you want to load database? y/n")

	db := new(database.Database)
	db.Users = make(map[uint64]database.User)

	if loadDatabase() {
		data, err := ioutil.ReadFile("database.json")
		if err != nil {
			log.Printf("%v\n", err)
		}
		db = &database.Database{}
		json.Unmarshal([]byte(data), &db)
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
