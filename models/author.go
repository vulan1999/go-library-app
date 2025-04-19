package models

import (
	"fmt"
	"os"
	"time"
)

type Author struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *Author) TableName() string {
	return fmt.Sprintf("%s.authors", os.Getenv("PG_SCHEMA"))
}
