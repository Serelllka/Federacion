package repository

import (
	"fmt"
	"github.com/Serelllka/Federacion/entities"
	"github.com/jmoiron/sqlx"
)

type ItemPostgres struct {
	db *sqlx.DB
}

func NewItemPostgres(db *sqlx.DB) *ItemPostgres {
	return &ItemPostgres{db: db}
}

func (r *ItemPostgres) CreateItem(item entities.Item) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, image, rarity) values ($1, $2, $3) RETURNING id", itemsTable)

	row := r.db.QueryRow(query, item.Title, item.Image, item.Rarity)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ItemPostgres) GetItemById(id int) (entities.Item, error) {
	var item entities.Item

	query := fmt.Sprintf("SELECT title, image, ratity FROM %s WHERE id = $1", itemsTable)
	err := r.db.Get(&item, query, id)

	return item, err
}

func (r *ItemPostgres) GetAllItems() (items []entities.Item, err error) {
	query := fmt.Sprintf("SELECT * FROM %s", itemsTable)
	err = r.db.Select(&items, query)

	return items, err
}
