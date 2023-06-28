package service

import (
	"github.com/Serelllka/Federacion/entities"
	"github.com/Serelllka/Federacion/internal/repository"
)

type CondemnationService struct {
	repo repository.Condemnation
}

func NewCondemnationService(repo repository.Condemnation) *CondemnationService {
	return &CondemnationService{repo: repo}
}

func (s *CondemnationService) CreateCondemnation(condemnation entities.Condemnation) (int, error) {
	return s.repo.CreateCondemnation(condemnation)
}

func (s *CondemnationService) GetAllCondemnations() ([]entities.Condemnation, error) {
	return s.repo.GetAllCondemnations()
}

func (s *CondemnationService) GetCondemnationById(id int) (entities.Condemnation, error) {
	return s.repo.GetCondemnationById(id)
}
