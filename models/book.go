package models

type Book struct {
	Id     uint   `json:"id"`
	Title  string `json:"title"`
	Author Author `json:"author"`
	Genre  Genre  `json:"genre"`
}
