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
	GetAllCondemnations() ([]entities.Condemnation, error)
}

type Users interface {
	GetAllUsers() ([]entities.UserDto, error)

	GetUserInfoById(id int) (entities.UserInfo, error)
}

type Achievements interface {
	GetAllAchievements() ([]entities.Achievement, error)
	GetAchievementById(id int) (entities.Achievement, error)
}

type UserAchievements interface {
	GetUserAchievements(id int) ([]entities.UserAchievement, error)
	GetUserAchievementsById(id int) (entities.UserAchievement, error)
}

type Repository struct {
	Achievements
	Authorization
	Article
	Users
	UserAchievements
	Condemnation
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:    NewAuthPostgres(db),
		Article:          NewArticlePostgres(db),
		Condemnation:     NewCondemnationPostgres(db),
		Users:            NewUsersPostgres(db),
		Achievements:     NewAchievementPostgres(db),
		UserAchievements: NewUserAchievementsPostgres(db),
	}
}
