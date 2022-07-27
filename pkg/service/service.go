package service

import (
	"github.com/Serelllka/Federacion/entities"
	"github.com/Serelllka/Federacion/pkg/repository"
)

type Authorization interface {
	CreateUser(user entities.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Articles interface {
	CreateArticle(article entities.Article) (int, error)
	GetAllArticles() ([]entities.Article, error)
	GetArticleById(id int) (entities.Article, error)
	UpdateArticle(id int, newArticle entities.Article) error
}

type Condemnations interface {
	CreateCondemnation(condemnation entities.Condemnation) (int, error)
	GetAllCondemnations() ([]entities.Condemnation, error)
	GetCondemnationById(id int) (entities.Condemnation, error)
}

type Users interface {
	GetAllUsers() (users []entities.UserInfo, err error)
}

type Service struct {
	Authorization
	Condemnations
	Articles
	Users
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Condemnations: NewCondemnationService(repos.Condemnation),
		Articles:      NewArticleService(repos.Article),
		Users:         NewUsersService(repos.Users),
	}
}
