package repository

import (
	"fmt"
	"github.com/Serelllka/Federacion/entities"
	"github.com/jmoiron/sqlx"
)

type CondemnationPostgres struct {
	db *sqlx.DB
}

func NewCondemnationPostgres(db *sqlx.DB) *CondemnationPostgres {
	return &CondemnationPostgres{db: db}
}

func (r *CondemnationPostgres) CreateCondemnation(condemnation entities.Condemnation) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (condemner_id, convicted_id, article_id, description, cost, time) "+
		"values ($1, $2, $3, $4, $5, $6) RETURNING id", condemnationsTable)

	row := r.db.QueryRow(query, condemnation.CondemnerId, condemnation.ConvictedId,
		condemnation.ArticleId, condemnation.Description, condemnation.Cost, condemnation.Time)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *CondemnationPostgres) GetCondemnationById(id int) (entities.Condemnation, error) {
	return entities.Condemnation{}, nil
}

func (r *CondemnationPostgres) UpdateCondemnation(id int, newCondemnations entities.Condemnation) error {
	return nil
}

func (r *CondemnationPostgres) GetAllCondemnations() (condemnations []entities.Condemnation, err error) {
	return []entities.Condemnation{}, nil
}
