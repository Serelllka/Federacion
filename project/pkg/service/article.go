package service

import (
	"github.com/Serelllka/Federacion/entities"
	"github.com/Serelllka/Federacion/pkg/repository"
)

type ArticleService struct {
	repo repository.Article
}

func NewArticleService(repo repository.Article) *ArticleService {
	return &ArticleService{repo: repo}
}

func (s *ArticleService) CreateArticle(article entities.Article) (int, error) {
	return s.repo.CreateArticle(article)
}

func (s *ArticleService) GetArticleById(id int) (entities.Article, error) {
	return s.repo.GetArticleById(id)
}

func (s *ArticleService) UpdateArticle(id int, newArticle entities.Article) error {
	return s.repo.UpdateArticle(id, newArticle)
}

func (s *ArticleService) GetAllArticles() ([]entities.Article, error) {
	return s.repo.GetAllArticles()
}
