package repository

import (
	"fmt"
	"github.com/Serelllka/Federacion/entities"
	"github.com/jmoiron/sqlx"
)

type UsersPostgres struct {
	db *sqlx.DB
}

func NewUsersPostgres(db *sqlx.DB) *UsersPostgres {
	return &UsersPostgres{db: db}
}

func (r *UsersPostgres) GetAllUsers() (users []entities.UserDto, err error) {
	query := fmt.Sprintf("SELECT id, name, username FROM %s", usersTable)
	err = r.db.Select(&users, query)

	return users, err
}
