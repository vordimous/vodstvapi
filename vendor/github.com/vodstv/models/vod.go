package models

import "time"

//Vod ...
type Vod struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
	MatchID   uint       `json:"-"`
	Tags      []Tag      `gorm:"many2many:vod_tags;"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	VideoKey  string     `json:"videoKey"`
	VideoURL  string     `json:"videoURL"`
	VideoSrc  string     `json:"videoSrc"`
	VideoDate string     `json:"videoDate"`
	ThumbURL  string     `json:"thumbURL"`
}
