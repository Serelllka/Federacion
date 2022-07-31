package entities

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type UserDto struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type UserInfo struct {
	Id        int    `json:"-" db:"id"`
	Name      string `json:"name" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Image     string `json:"image"`
	Hobbies   string `json:"hobbies"`
	Lore      string `json:"lore"`
	Sobriquet string `json:"sobriquet"`
}
