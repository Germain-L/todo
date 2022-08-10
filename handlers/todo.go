package handlers

import (
	"encoding/json"
	"gotodo/models"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
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

func (h Handler) PatchTodo(w http.ResponseWriter, r *http.Request) {
	// get email from context
	var email = r.Context().Value("email").(string)

	var user models.User

	// get userid from email
	user, err := h.UserRepo.GetUserEmail(email)
	if err != nil {
		log.Println("error getting user")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	// get todo id from mux vars
	vars := mux.Vars(r)
	var todoId = vars["id"]

	// get todo from id
	todo, err := h.TodoRepo.GetTodoId(todoId)
	if err != nil {
		log.Println("error getting todo")
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	// check if todo belongs to user
	if todo.User_id != user.Id {
		log.Println("todo does not belong to user")
		log.Println(todo.User_id)
		log.Println(user.Id)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	type Completed struct {
		Completed bool `json:"completed"`
	}

	var completed Completed
	err = json.NewDecoder(r.Body).Decode(&completed)
	if err != nil {
		log.Println("error decoding completed")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	if completed.Completed {
		err = h.TodoRepo.Complete(todoId)
		if err != nil {
			log.Println("error completing todo")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}
	} else {
		err = h.TodoRepo.Incomplete(todoId)
		if err != nil {
			log.Println("error completing todo")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func (h Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	// get email from context
	var email = r.Context().Value("email").(string)

	var user models.User

	// get userid from email
	user, err := h.UserRepo.GetUserEmail(email)
	if err != nil {
		log.Println("error getting user")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	// get todo id from mux vars
	vars := mux.Vars(r)
	var todoId = vars["id"]

	// get todo from id
	todo, err := h.TodoRepo.GetTodoId(todoId)
	if err != nil {
		log.Println("error getting todo")
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	// check if todo belongs to user
	if todo.User_id != user.Id {
		log.Println("todo does not belong to user")
		log.Println(todo.User_id)
		log.Println(user.Id)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	err = h.TodoRepo.DeleteTodo(todo.Id)
	if err != nil {
		log.Println("error deleting todo")
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
}
