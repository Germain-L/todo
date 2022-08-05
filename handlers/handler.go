package handlers

import (
	"gotodo/repository"
)

type Handler struct {
	UserRepo repository.UserRepo
	TodoRepo repository.TodoRepository
}

func New(userRepo repository.UserRepo, todoRepository repository.TodoRepository) Handler {
	return Handler{
		UserRepo: userRepo,
		TodoRepo: todoRepository,
	}
}
