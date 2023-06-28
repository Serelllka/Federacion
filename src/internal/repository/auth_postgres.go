package repository

import (
	"fmt"
	"github.com/Serelllka/Federacion/entities"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (ap *AuthPostgres) CreateUser(user entities.User) (int, error) {
	tx, err := ap.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createUser := fmt.Sprintf(
		"INSERT INTO %s (name, username, email, password_hash) values ($1, $2, $3, $4) RETURNING id", usersTable)
	row := tx.QueryRow(createUser, user.Name, user.Username, user.Email, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	createUserInfo := fmt.Sprintf(
		"INSERT INTO %s (user_id) values ($1) RETURNING id", usersInfoTable)
	_, err = tx.Exec(createUserInfo, id)
	if err != nil {
		return 0, err
	}

	return id, tx.Commit()
}

func (ap *AuthPostgres) GetUser(username, password string) (entities.User, error) {
	var user entities.User
	query := fmt.Sprintf(
		"SELECT id from %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := ap.db.Get(&user, query, username, password)

	return user, err
}
