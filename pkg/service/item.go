package service

import (
	"github.com/Serelllka/Federacion/entities"
	"github.com/Serelllka/Federacion/pkg/repository"
)

type ItemService struct {
	repo repository.Items
}

func NewItemService(repo repository.Items) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) CreateItem(item entities.Item) (int, error) {
	return s.repo.CreateItem(item)
}

func (s *ItemService) GetItemById(id int) (entities.Item, error) {
	return s.repo.GetItemById(id)
}

func (s *ItemService) GetAllItems() ([]entities.Item, error) {
	return s.repo.GetAllItems()
}
