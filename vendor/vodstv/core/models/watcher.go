package models

import (
	"time"

	"vodstv/core"

	"golang.org/x/crypto/bcrypt"
)

//Watcher ...
type Watcher struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"_"`
	UpdatedAt time.Time  `json:"_"`
	DeletedAt *time.Time `json:"_" sql:"index"`
	IsAdmin   bool       `json:"isAdmin"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	Name      string     `json:"name"`
	Feeds     []Feed     `gorm:"ForeignKey:OwnerID"`
}

//BeforeCreate ...
func (w *Watcher) BeforeCreate() (err error) {
	bytePassword := []byte(w.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	core.CheckErr(err, "Pass hash failed")

	w.Password = string(hashedPassword)
	return
}
