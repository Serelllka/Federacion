package repository

import (
	"github.com/Serelllka/Federacion/entities"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user entities.User) (int, error)
	GetUser(username, password string) (entities.User, error)
}

type Article interface {
	CreateArticle(article entities.Article) (int, error)
	GetArticleById(id int) (entities.Article, error)
	UpdateArticle(id int, newArticle entities.Article) error
	GetAllArticles() ([]entities.Article, error)
}

type Condemnation interface {
	CreateCondemnation(condemnations entities.Condemnation) (int, error)
	GetCondemnationById(id int) (entities.Condemnation, error)
	UpdateCondemnation(id int, newCondemnations entities.Condemnation) error
	GetAllCondemnations() (condemnations []entities.Condemnation, err error)
}

type Users interface {
	GetAllUsers() (users []entities.UserDto, err error)

	GetUserInfoById(id int) (userInfo entities.UserInfo, err error)
}

type Achievements interface {
	GetAllAchievements() (achievements []entities.Achievement)
}

type Repository struct {
	Users
	Condemnation
	Authorization
	Article
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Article:       NewArticlePostgres(db),
		Condemnation:  NewCondemnationPostgres(db),
		Users:         NewUsersPostgres(db),
	}
}
