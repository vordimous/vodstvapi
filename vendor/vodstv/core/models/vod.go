package models

import "time"

//Vod ...
type Vod struct {
	ID          uint       `json:"id" gorm:"primary_key"`
	CreatedAt   time.Time  `json:"_"`
	UpdatedAt   time.Time  `json:"_"`
	DeletedAt   *time.Time `json:"_" sql:"index"`
	MatchID     uint       `json:"-"`
	Tags        []Tag      `json:"tags" gorm:"many2many:vod_tags;"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	VideoKey    string     `json:"videoKey" gorm:"not null;unique"`
	VideoURL    string     `json:"videoURL"`
	VideoSrc    string     `json:"videoSrc"`
	VideoDate   time.Time  `json:"videoDate"`
	ThumbURL    string     `json:"thumbURL"`
}
