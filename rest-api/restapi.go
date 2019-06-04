package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/GrzegorzCzaprowski/go-exercises/rest-api/handlers"
	"github.com/GrzegorzCzaprowski/go-exercises/rest-api/models"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "postgres://testuser:testpass@localhost:5555/testdb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()

	server := models.Server{DB: db}

	h := handlers.Handler{
		S: server}

	router.HandleFunc("/api/todos/", h.Post).Methods("POST")
	router.HandleFunc("/api/todos/", h.GetAll).Methods("GET")
	// router.HandleFunc("/api/todos/:id/", ReadByID).Methods("GET")
	// router.HandleFunc("/api/todos/:id/", UpdateTodo).Methods("PATCH")
	// router.HandleFunc("/api/todos/:id/", RemoveTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}
