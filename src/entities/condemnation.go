package entities

import "time"

type Condemnation struct {
	Id          int       `json:"-" db:"id"`
	CondemnerId int       `json:"-" db:"condemner_id"`
	ConvictedId int       `json:"convicted_id" db:"convicted_id" binding:"required"`
	ArticleId   int       `json:"article_id" db:"article_id" binding:"required"`
	Description string    `json:"description" db:"description"`
	Cost        int       `json:"cost" db:"cost"`
	Time        time.Time `json:"-" db:"time"`
}
