package service

import (
	federacion "github.com/Serelllka/Federacion/entities"
	"github.com/Serelllka/Federacion/pkg/repository"
)

type ArticleService struct {
	repo repository.Article
}

func NewArticleService(repo repository.Article) *ArticleService {
	return &ArticleService{repo: repo}
}

func (s *ArticleService) CreateArticle(article federacion.Article) (int, error) {
	return s.repo.CreateArticle(article)
}

func (s *ArticleService) GetArticleById(id int) (federacion.Article, error) {
	return s.repo.GetArticleById(id)
}

func (s *ArticleService) UpdateArticle(id int, newArticle federacion.Article) error {
	return s.repo.UpdateArticle(id, newArticle)
}

func (s *ArticleService) GetAllArticles() ([]federacion.Article, error) {
	return s.repo.GetAllArticles()
}
