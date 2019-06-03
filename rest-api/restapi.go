//https://medium.com/@adigunhammedolalekan/build-and-deploy-a-secure-rest-api-with-go-postgresql-jwt-and-gorm-6fadf3da505b
//https://golang.org/doc/effective_go.html

package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/GrzegorzCzaprowski/go-exercises/rest-api/models"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "postgres://testuser:testpass@localhost:5555/testdb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server := models.Server{DB: db}

	router := mux.NewRouter()

	router.HandleFunc("/api/todos/", server.CreateTodo).Methods("POST")
	router.HandleFunc("/api/todos/", server.ReadAllTodos).Methods("GET")
	// router.HandleFunc("/api/todos/:id/", ReadByID).Methods("GET")
	// router.HandleFunc("/api/todos/:id/", UpdateTodo).Methods("PATCH")
	// router.HandleFunc("/api/todos/:id/", RemoveTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
