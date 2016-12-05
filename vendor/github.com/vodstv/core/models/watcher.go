package models

import (
	"time"

	"github.com/vodstv/core"

	"golang.org/x/crypto/bcrypt"
)

//Watcher ...
type Watcher struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
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
