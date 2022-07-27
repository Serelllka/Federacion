package service

import (
	federacion "github.com/Serelllka/Federacion/entities"
	"github.com/Serelllka/Federacion/pkg/repository"
)

type CondemnationService struct {
	repo repository.Condemnation
}

func NewCondemnationService(repo repository.Condemnation) *CondemnationService {
	return &CondemnationService{repo: repo}
}

func (s *CondemnationService) CreateCondemnation(condemnation federacion.Condemnation) (int, error) {
	return s.repo.CreateCondemnation(condemnation)
}

func (s *CondemnationService) GetAllCondemnations() ([]federacion.Condemnation, error) {
	return []federacion.Condemnation{}, nil
}

func (s *CondemnationService) GetCondemnationById(id int) (federacion.Condemnation, error) {
	return federacion.Condemnation{}, nil
}
