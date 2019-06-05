package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	"github.com/GrzegorzCzaprowski/go-exercises/rest-api/handlers"
	"github.com/GrzegorzCzaprowski/go-exercises/rest-api/models"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {
	var flagDatabaseAddress string
	flag.StringVar(&flagDatabaseAddress, "db", "postgres://testuser:testpass@localhost:5555/testdb?sslmode=disable", "database address")
	flag.Parse()

	db, err := sql.Open("postgres", flagDatabaseAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := httprouter.New()

	mod := models.Model{DB: db}
	handler := handlers.Handler{M: mod}

	router.POST("/api/todos/", handler.Post)
	router.GET("/api/todos/", handler.GetAll)
	router.GET("/api/todos/:id/", handler.Get)
	router.PATCH("/api/todos/:id/", handler.Patch)
	router.DELETE("/api/todos/:id/", handler.Delete)

	log.Fatal(http.ListenAndServe(":8000", router))
}
