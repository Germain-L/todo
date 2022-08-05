package repository

import "database/sql"

type TodoRepository struct {
	Db *sql.DB
}
