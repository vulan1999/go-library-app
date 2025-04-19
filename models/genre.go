package models

import (
	"fmt"
	"os"
)

type Genre struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func (g *Genre) TableName() string {
	return fmt.Sprintf("%s.genres", os.Getenv("PG_SCHEMA"))
}
