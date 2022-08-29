package entities

type Item struct {
	Id     int    `json:"id" db:"id"`
	Title  string `json:"title" db:"title" binding:"required"`
	Image  string `json:"image" db:"image"`
	Rarity int    `json:"rarity" db:"image"`
}
