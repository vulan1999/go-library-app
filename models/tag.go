package models

import (
	"fmt"
	"os"
	"time"
)

type Tag struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

func (t *Tag) TableName() string {
	return fmt.Sprintf("%s.tags", os.Getenv("PG_SCHEMA"))
}
