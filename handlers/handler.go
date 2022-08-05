package handlers

import (
	"gotodo/repository"
)

type Handler struct {
	UserRepo repository.UserRepo
	TodoRepo repository.TodoRepo
}

func New(userRepo repository.UserRepo, todoRepository repository.TodoRepo) Handler {
	return Handler{
		UserRepo: userRepo,
		TodoRepo: todoRepository,
	}
}
