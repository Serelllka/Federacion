package entities

import "time"

type UserAchievement = struct {
	Id            int       `json:"id" db:"id"`
	UserId        int       `json:"userId" db:"user_id"`
	AchievementId int       `json:"achievementId" db:"achievement_id"`
	Name          string    `json:"name" db:"name"`
	Description   string    `json:"description" db:"description"`
	Image         string    `json:"image" db:"image"`
	Time          time.Time `json:"-" db:"time"`
}
