package service

import (
	"github.com/Serelllka/Federacion/entities"
	"github.com/Serelllka/Federacion/internal/repository"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{repo: repo}
}

func (s *UsersService) GetAllUsers() (users []entities.UserDto, err error) {
	return s.repo.GetAllUsers()
}

func (s *UsersService) GetUserInfoById(id int) (userInfo entities.UserInfo, err error) {
	return s.repo.GetUserInfoById(id)
}
