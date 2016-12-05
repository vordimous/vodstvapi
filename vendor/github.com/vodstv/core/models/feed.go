package models

import "time"

//Feed ...
type Feed struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
	OwnerID   uint       `json:"-"`
	Name      string     `json:"name"`
	Tags      []Tag      `gorm:"many2many:feed_tags;"`
}
