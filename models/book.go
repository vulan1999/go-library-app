package models

import "time"

type Book struct {
	Id         uint      `json:"id"`
	Title      string    `json:"title"`
	Author     Author    `json:"author"`
	Genres     []Genre   `json:"genre"`
	Tags       []Tag     `json:"tags"`
	Language   Language  `json:"language"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
