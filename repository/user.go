package repository

import (
	"database/sql"
	"gotodo/models"
)

type UserRepo struct {
	Db *sql.DB
}

func (repo UserRepo) GetUserId(id string) (models.User, error) {
	var user models.User

	q, err := repo.Db.Query("SELECT * FROM users WHERE id = $1", id)

	if err != nil {
		return user, err
	}

	q.Scan(user)

	return user, nil
}

func (repo UserRepo) GetUserEmail(email string) (models.User, error) {
	var user models.User

	// query the database for the user with the given email
	// put it in user
	err := repo.Db.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo UserRepo) CreateUser(user models.User) error {
	_, err := repo.Db.Exec("INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)", user.Id, user.Name, user.Email, user.Password)

	return err
}

func (repo UserRepo) Deleteuser(id string) error {
	_, err := repo.Db.Exec("DELETE FROM users WHERE id = $1", id)

	return err
}
