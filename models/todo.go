package models

import "time"

type Todo struct {
	Id         string    `json:"id"`
	Title      string    `json:"title"`
	Completed  bool      `json:"completed"`
	User_id    string    `json:"user_id"`
	Created_at time.Time `json:"created_at"`
}
