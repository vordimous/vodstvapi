package models

import "time"

//Feed ...
type Feed struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"_"`
	UpdatedAt time.Time  `json:"_"`
	DeletedAt *time.Time `json:"_" sql:"index"`
	OwnerID   uint       `json:"-"`
	Name      string     `json:"name"`
	Tags      []Tag      `json:"tags" gorm:"many2many:feed_tags;"`
}
