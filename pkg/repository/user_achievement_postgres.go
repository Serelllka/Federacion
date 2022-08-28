package repository

import (
	"fmt"
	"github.com/Serelllka/Federacion/entities"
	"github.com/jmoiron/sqlx"
)

type UserAchievementsPostgres struct {
	db *sqlx.DB
}

func NewUserAchievementsPostgres(db *sqlx.DB) *UserAchievementsPostgres {
	return &UserAchievementsPostgres{db: db}
}

func (r *UserAchievementsPostgres) GetUserAchievements(id int) (userAchievements []entities.UserAchievement, err error) {
	query := fmt.Sprintf("SELECT  ua.id, user_id, achievement_id, time, name FROM %s AS ua "+
		"JOIN %s AS a ON ua.achievement_id = a.id WHERE ua.user_id = $1",
		usersAchievementsTable, achievementsTable)
	err = r.db.Select(&userAchievements, query, id)

	return userAchievements, err
}

func (r *UserAchievementsPostgres) GetUserAchievementsById(id int) (userAchievement entities.UserAchievement, err error) {
	return entities.UserAchievement{}, err
}
