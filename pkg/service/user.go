package service

import (
	federacion "github.com/Serelllka/Federacion/entities"
	"github.com/Serelllka/Federacion/pkg/repository"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{repo: repo}
}

func (s *UsersService) GetAllUsers() (users []federacion.UserInfo, err error) {
	return s.repo.GetAllUsers()
}
