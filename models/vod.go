package models

import "github.com/jinzhu/gorm"

//Vod ...
type Vod struct {
	gorm.Model
	Email   string `json:"email"`
	Name    string `json:"name"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
