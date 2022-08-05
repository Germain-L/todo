package handlers

import (
	"encoding/json"
	"gotodo/models"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (h Handler) PostTodo(w http.ResponseWriter, r *http.Request) {
	var email = r.Context().Value("email").(string)

	// decode todo from body
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	// get userid from email
	user, err := h.UserRepo.GetUserEmail(email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	// set userid to todo
	todo.User_id = user.Id

	// generate id for todo
	todo.Id = uuid.NewString()

	// set created at to now
	todo.Created_at = time.Now()

	// create todo
	err = h.TodoRepo.CreateTodo(todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func (h Handler) GetTodo(w http.ResponseWriter, r *http.Request) {
	var email = r.Context().Value("email").(string)

	// get userid from email
	user, err := h.UserRepo.GetUserEmail(email)
	if err != nil {
		log.Println("error getting user")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	// get todos from userid
	todos, err := h.TodoRepo.GetTodosUser(user.Id)
	if err != nil {
		log.Println("error getting todos")
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}
