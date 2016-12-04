package models

import "time"

//Tag ...
type Tag struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
	Name      string     `json:"name"`
	Type      string     `json:"type"`
	Regex     string     `json:"regex"`
}
