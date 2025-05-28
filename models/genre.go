package models

import (
	"time"
)

type Genre struct {
	Id          uint      `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (g *Genre) TableName() string {
	return "genres"
}
