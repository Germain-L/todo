package repository

import (
	"database/sql"
	"gotodo/models"
)

type TodoRepo struct {
	Db *sql.DB
}

func (repo TodoRepo) GetTodosUser(user_id string) ([]models.Todo, error) {
	var todos []models.Todo

	rows, err := repo.Db.Query("SELECT * FROM todo WHERE user_id = $1", user_id)
	if err != nil {
		return todos, err
	}

	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(
			&todo.Id,
			&todo.Title,
			&todo.Completed,
			&todo.Created_at,
			&todo.User_id,
		)

		if err != nil {
			return todos, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (repo TodoRepo) GetTodoId(id string) (models.Todo, error) {
	var todo models.Todo

	err := repo.Db.QueryRow("SELECT * FROM todo WHERE id = $1", id).Scan(
		&todo.Id,
		&todo.Title,
		&todo.Completed,
		&todo.Created_at,
		&todo.User_id,
	)

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (repo TodoRepo) GetTodoTitle(email string) (models.Todo, error) {
	var todo models.Todo

	// query the database for the user with the given email
	// put it in user
	err := repo.Db.QueryRow("SELECT * FROM todo WHERE title = $1", email).Scan(
		&todo.Id,
		&todo.Title,
		&todo.Completed,
		&todo.Created_at,
		&todo.User_id,
	)

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (repo TodoRepo) CreateTodo(todo models.Todo) error {
	_, err := repo.Db.Exec("INSERT INTO todo (id, title, completed, user_id) VALUES ($1, $2, $3, $4)",
		todo.Id,
		todo.Title,
		todo.Completed,
		todo.Created_at,
		todo.User_id,
	)

	return err
}

func (repo TodoRepo) Complete(id string) error {
	_, err := repo.Db.Exec("UPDATE todo SET completed = true WHERE id = $1", id)

	return err
}

func (repo TodoRepo) Incomplete(id string) error {
	_, err := repo.Db.Exec("UPDATE todo SET completed = false WHERE id = $1", id)

	return err
}

func (repo TodoRepo) DeleteTodo(id string) error {
	_, err := repo.Db.Exec("DELETE FROM todo WHERE id = $1", id)

	return err
}
