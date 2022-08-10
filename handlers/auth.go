package handlers

import (
	"encoding/json"
	"gotodo/auth"
	"gotodo/models"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// Read to request body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var user models.User
	json.Unmarshal(body, &user)

	hashedPass, err := auth.HashPassword(user.Password)
	if err != nil {
		log.Println("error hashing :", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.Id = uuid.NewString()
	user.Password = hashedPass

	log.Println(user)

	err = h.UserRepo.CreateUser(user)
	if err != nil {
		log.Println(err)

		if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode("User already exists")
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	// Send a 201 created response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func (h Handler) Signin(w http.ResponseWriter, r *http.Request) {
	// Get the email and password from basic auth in request
	email, password, ok := r.BasicAuth()

	var user models.User

	// get password of user in database
	user, err := h.UserRepo.GetUserEmail(email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No users found"))
		return
	}

	if !auth.CheckPasswordHash(password, user.Password) || !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("No/Invalid Credentials"))
		return
	}

	// get new token, add user.Email to it
	tokenString, expirationTime, err := auth.GetNewToken(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// write cookie to response
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	log.Printf("Success, %s issued", tokenString)
	json.NewEncoder(w).Encode(map[string]string{"id": user.Id})
}

func (h Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var email = r.Context().Value("email").(string)

	// get userid from email
	user, err := h.UserRepo.GetUserEmail(email)
	if err != nil {
		log.Println("error getting user")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	var todos []models.Todo
	todos, err = h.TodoRepo.GetTodosUser(user.Id)
	if err != nil {
		log.Println("error deleting user's todos")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	for _, v := range todos {
		err = h.TodoRepo.DeleteTodo(v.Id)
		if err != nil {
			log.Println("error deleting user's todos")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}
	}

	err = h.UserRepo.Deleteuser(user.Id)
	if err != nil {
		log.Println("error deleting user")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
}
