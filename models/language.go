package models

type Language struct {
	Id         int    `json:"id"`
	Language   string `json:"language"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
