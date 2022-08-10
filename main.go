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
	tr := repository.TodoRepo{Db: db}

	h := handlers.New(ur, tr)

	r.HandleFunc("/register", h.Register).Methods(http.MethodPost)
	r.HandleFunc("/signin", h.Signin).Methods(http.MethodGet)

	r.Handle("/todo", handlers.Authenticate(http.HandlerFunc(h.PostTodo))).Methods(http.MethodPost)
	r.Handle("/todo", handlers.Authenticate(http.HandlerFunc(h.GetTodo))).Methods(http.MethodGet)
	r.Handle("/todo/{id}", handlers.Authenticate(http.HandlerFunc(h.PatchTodo))).Methods(http.MethodPatch)
	r.Handle("/todo/{id}", handlers.Authenticate(http.HandlerFunc(h.DeleteTodo))).Methods(http.MethodDelete)

	r.Handle("/user/{id}", handlers.Authenticate(http.HandlerFunc(h.DeleteUser))).Methods(http.MethodDelete)

	return r
}

func main() {
	loadEnv()
	postgres := db.ConnectPostgres()

	server := setupServer(postgres)

	log.Println("http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", server))
}
