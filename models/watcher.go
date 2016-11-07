package models

import "github.com/jinzhu/gorm"

//Watcher ...
type Watcher struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"-"`
	Name     string `json:"name"`
}
