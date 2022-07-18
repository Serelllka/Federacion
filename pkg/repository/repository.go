package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Friends interface {
}

type Repository struct {
	Authorization
	Friends
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
