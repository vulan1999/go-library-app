package models

import (
	"time"
)

type Tag struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Books       []*Book   `gorm:"many2many:book_tags"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

func (t *Tag) TableName() string {
	return "tags"
}

type BookTag struct {
	Id         uint
	BookId     uint
	TagId      uint
	Created_at time.Time
	Updated_at time.Time
}
