package entities

type Article struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name" binding:"required"`
	CurrentCost string `json:"currentCost" db:"current_cost" binding:"required"`
	Description string `json:"description" db:"description"`
}
