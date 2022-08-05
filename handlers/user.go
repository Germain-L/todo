package handlers

import (
	"encoding/json"
	"gotodo/models"
	"net/http"

	"github.com/gorilla/mux"
)

func (h Handler) GetUserId(w http.ResponseWriter, r *http.Request) {
	var user models.User

	vars := mux.Vars(r)
	id := vars["id"]

	user, err := h.UserRepo.GetUserId(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
