package models

import (
	"fmt"
	"os"
	"time"
)

type Book struct {
	Id             uint      `json:"id"`
	Title          string    `json:"title"`
	AuthorId       uint      `json:"-"`
	Author         Author    `json:"author" gorm:"foreignKey:AuthorId;reference:Id"`
	LanguageId     uint      `json:"-"`
	Language       Language  `json:"language" gorm:"foreignKey:LanguageId;reference:Id"`
	OriginalBookId *uint     `json:"-"`
	OriginalBook   *Book     `json:"original_book" gorm:"foreignKey:OriginalBookId;reference:Id"`
	Created_at     time.Time `json:"created_at"`
	Updated_at     time.Time `json:"updated_at"`
}

func (b *Book) TableName() string {
	return fmt.Sprintf("%s.books", os.Getenv("PG_SCHEMA"))
}
