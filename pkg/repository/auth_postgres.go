package repository

import (
	"fmt"
	federacion "github.com/Serelllka/Federacion"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (ap *AuthPostgres) CreateUser(user federacion.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNINNG id")

	row := ap.db.QueryRow(query, user.Name, user.Id, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
