package models

import "time"

//Tag ...
type Tag struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"_"`
	UpdatedAt time.Time  `json:"_"`
	DeletedAt *time.Time `json:"_" sql:"index"`
	Name      string     `json:"name"`
	Type      string     `json:"type"`
	Regex     string     `json:"regex"`
}
