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

func (r *UsersPostgres) GetUserInfoById(id int) (userInfo entities.UserInfo, err error) {
	query := fmt.Sprintf(
		`SELECT u.id, u.name, u.username, u.email, ui.image, ui.hobbies, ui.lore, 
	ui.sobriquet from %s as u join %s as ui on u.id=ui.user_id where u.id=$1`, usersTable, usersInfoTable)
	err = r.db.Get(&userInfo, query, id)

	return userInfo, err
}
