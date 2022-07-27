package repository

import (
	"fmt"
	federacion "github.com/Serelllka/Federacion/entities"
	"github.com/jmoiron/sqlx"
)

type UsersPostgres struct {
	db *sqlx.DB
}

func NewUsersPostgres(db *sqlx.DB) *UsersPostgres {
	return &UsersPostgres{db: db}
}

func (r *UsersPostgres) GetAllUsers() (users []federacion.UserInfo, err error) {
	query := fmt.Sprintf("SELECT id, name, username FROM %s", usersTable)
	err = r.db.Select(&users, query)

	return users, err
}
