package repository

import (
	federacion "github.com/Serelllka/Federacion"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user federacion.User) (int, error)
	GetUser(username, password string) (federacion.User, error)
}

type Friends interface {
}

type Repository struct {
	Authorization
	Friends
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
