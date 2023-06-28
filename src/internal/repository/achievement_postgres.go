package repository

import (
	"fmt"
	"github.com/Serelllka/Federacion/entities"
	"github.com/jmoiron/sqlx"
)

type AchievementPostgres struct {
	db *sqlx.DB
}

func NewAchievementPostgres(db *sqlx.DB) *AchievementPostgres {
	return &AchievementPostgres{db: db}
}

func (r *AchievementPostgres) GetAllAchievements() (achievements []entities.Achievement, err error) {
	query := fmt.Sprintf("SELECT * FROM %s", achievementsTable)
	err = r.db.Select(&achievements, query)

	return achievements, err
}

func (r *AchievementPostgres) GetAchievementById(id int) (entities.Achievement, error) {
	var achievement entities.Achievement

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", achievementsTable)
	err := r.db.Get(&achievement, query, id)

	return achievement, err
}
