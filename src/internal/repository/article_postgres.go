package repository

import (
	"fmt"
	"github.com/Serelllka/Federacion/entities"
	"github.com/jmoiron/sqlx"
)

type ArticlePostgres struct {
	db *sqlx.DB
}

func NewArticlePostgres(db *sqlx.DB) *ArticlePostgres {
	return &ArticlePostgres{db: db}
}

func (r *ArticlePostgres) CreateArticle(article entities.Article) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, current_cost, description) values ($1, $2, $3) RETURNING id", articlesTable)

	row := r.db.QueryRow(query, article.Name, article.CurrentCost, article.Description)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ArticlePostgres) GetArticleById(id int) (entities.Article, error) {
	var article entities.Article

	query := fmt.Sprintf("SELECT name, current_cost, description FROM %s WHERE id = $1", articlesTable)
	err := r.db.Get(&article, query, id)

	return article, err
}

func (r *ArticlePostgres) UpdateArticle(id int, newArticle entities.Article) error {
	return nil
}

func (r *ArticlePostgres) GetAllArticles() (articles []entities.Article, err error) {
	query := fmt.Sprintf("SELECT * FROM %s", articlesTable)
	err = r.db.Select(&articles, query)

	return articles, err
}
