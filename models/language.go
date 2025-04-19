package models

import (
	"fmt"
	"os"
)

type Language struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}

func (l *Language) TableName() string {
	return fmt.Sprintf("%s.languages", os.Getenv("PG_SCHEMA"))
}
