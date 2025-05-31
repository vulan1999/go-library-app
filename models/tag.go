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
	TagId          uint   `json:"tag_id"`
	TagDescription string `json:"tag_description"`
}
