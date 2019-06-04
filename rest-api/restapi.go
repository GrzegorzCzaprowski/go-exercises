package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/GrzegorzCzaprowski/go-exercises/rest-api/handlers"
	"github.com/GrzegorzCzaprowski/go-exercises/rest-api/models"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "postgres://testuser:testpass@localhost:5555/testdb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := httprouter.New()

	server := models.Model{DB: db}

	h := handlers.Handler{
		M: server}

	router.POST("/api/todos/", h.Post)
	router.GET("/api/todos/", h.GetAll)
	// router.HandleFunc("/api/todos/:id/", ReadByID).Methods("GET")
	// router.HandleFunc("/api/todos/:id/", UpdateTodo).Methods("PATCH")
	// router.HandleFunc("/api/todos/:id/", RemoveTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}
