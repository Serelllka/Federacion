package service

import (
	"github.com/Serelllka/Federacion/entities"
	"github.com/Serelllka/Federacion/pkg/repository"
)

type AchievementService struct {
	achievementsRepo     repository.Achievements
	userAchievementsRepo repository.UserAchievements
}

func NewAchievementService(
	achieve repository.Achievements,
	userAchieve repository.UserAchievements) *AchievementService {
	return &AchievementService{
		achievementsRepo:     achieve,
		userAchievementsRepo: userAchieve,
	}
}

func (s *AchievementService) GetAllAchievements() ([]entities.Achievement, error) {
	return s.achievementsRepo.GetAllAchievements()
}

func (s *AchievementService) GetAchievementById(id int) (entities.Achievement, error) {
	return s.achievementsRepo.GetAchievementById(id)
}

func (s *AchievementService) GetUserAchievements(id int) ([]entities.UserAchievement, error) {
	return s.userAchievementsRepo.GetUserAchievements(id)
}
