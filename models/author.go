package models

import (
	"time"
)

type Author struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *Author) TableName() string {
	return "authors"
}

type AuthorCreateRequest struct {
	Name string `json:"name" binding:"required"`
}
