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
	json.NewEncoder(w).Encode("Success")
}
