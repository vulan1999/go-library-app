package models

import (
	"encoding/json"
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
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (b *Book) TableName() string {
	return fmt.Sprintf("%s.books", os.Getenv("PG_SCHEMA"))
}

func (b *Book) MarshalJSON() ([]byte, error) {
	type BookAuthor struct {
		AuthorId   uint   `json:"author_id"`
		AuthorName string `json:"author_name"`
	}

	type BookLanguage struct {
		LanguageId          uint   `json:"language_id"`
		LanguageDescription string `json:"language_description"`
	}

	type OriginalBook struct {
		OriginalBookId    uint   `json:"original_book_id"`
		OriginalBookTitle string `json:"original_book_title"`
	}

	type Temp struct {
		Id           uint         `json:"id"`
		Title        string       `json:"title"`
		Author       BookAuthor   `json:"author"`
		Language     BookLanguage `json:"language"`
		OriginalBook OriginalBook `json:"original_book"`
		CreatedAt    time.Time    `json:"created_at"`
		UpdatedAt    time.Time    `json:"updated_at"`
	}

	t := Temp{
		Id:           b.Id,
		Title:        b.Title,
		Author:       BookAuthor{AuthorId: b.Author.Id, AuthorName: b.Author.Name},
		Language:     BookLanguage{LanguageId: b.Language.Id, LanguageDescription: b.Language.Description},
		OriginalBook: OriginalBook{OriginalBookId: b.OriginalBook.Id, OriginalBookTitle: b.OriginalBook.Title},
		CreatedAt:    b.CreatedAt,
		UpdatedAt:    b.UpdatedAt,
	}

	return json.Marshal(&t)
}
