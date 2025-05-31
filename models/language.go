package models

type Language struct {
	Id          uint   `json:"id"`
	Description string `json:"description"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}

func (l *Language) TableName() string {
	return "languages"
}

type LanguageResponse struct {
	LanguageId          uint   `json:"language_id"`
	LanguageDescription string `json:"language_description"`
}
