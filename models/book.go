package models

import (
	"encoding/json"
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
	Tags           []*Tag    `json:"tags" gorm:"many2many:book_tags"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type BookCreateRequest struct {
	Title          string `json:"title"`
	AuthorId       uint   `json:"author_id"`
	LanguageId     uint   `json:"language_id"`
	OriginalBookId *uint  `json:"original_book_id"`
}

func (b *Book) TableName() string {
	return "books"
}

type OriginalBook struct {
	OriginalBookId    uint   `json:"original_book_id,omitempty"`
	OriginalBookTitle string `json:"original_book_title,omitempty"`
}

type BookResponse struct {
	Id           uint             `json:"id"`
	Title        string           `json:"title"`
	Author       AuthorResponse   `json:"author"`
	Language     LanguageResponse `json:"language"`
	OriginalBook *OriginalBook    `json:"original_book"`
	Tags         []*BookTag       `json:"tags"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    time.Time        `json:"updated_at"`
}

func (b *Book) MarshalJSON() ([]byte, error) {
	var original_book OriginalBook

	if b.OriginalBook != nil {
		original_book = OriginalBook{OriginalBookId: b.OriginalBook.Id, OriginalBookTitle: b.OriginalBook.Title}
	}

	var tag_list []*BookTag
	for _, value := range b.Tags {
		tag_list = append(tag_list, &BookTag{TagId: uint(value.Id), TagDescription: value.Description})
	}

	t := BookResponse{
		Id:           b.Id,
		Title:        b.Title,
		Author:       AuthorResponse{AuthorId: b.Author.Id, AuthorName: b.Author.Name},
		Language:     LanguageResponse{LanguageId: b.Language.Id, LanguageDescription: b.Language.Description},
		OriginalBook: &original_book,
		Tags:         tag_list,
		CreatedAt:    b.CreatedAt,
		UpdatedAt:    b.UpdatedAt,
	}

	return json.Marshal(&t)
}
