package main

import (
	"database/sql"
	"gotodo/db"
	"gotodo/handlers"
	"gotodo/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Id := uuid.New()

func loadEnv() {
	err := godotenv.Load("env/backend.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func setupServer(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	ur := repository.UserRepo{Db: db}
	tr := repository.TodoRepository{Db: db}

	h := handlers.New(ur, tr)

	r.HandleFunc("/register", h.Register).Methods(http.MethodPost)
	r.HandleFunc("/signin", h.Signin).Methods(http.MethodGet)
	r.HandleFunc("/user/id", h.GetUserId).Methods(http.MethodGet)

	return r
}

func main() {
	loadEnv()
	postgres := db.ConnectPostgres()

	server := setupServer(postgres)

	log.Println("http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", server))
}
