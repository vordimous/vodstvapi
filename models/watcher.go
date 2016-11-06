package models

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"

	"esvodsApi/db"
	"esvodsApi/forms"

	"github.com/jinzhu/gorm"
)

//Watcher ...
type Watcher struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"-"`
	Name     string `json:"name"`
}

//WatcherModel ...
type WatcherModel struct{}

//Signin ...
func (m WatcherModel) Signin(form forms.SigninForm) (watcher Watcher, err error) {
	db.GetDB().Where(&Watcher{Email: form.Email}).First(&watcher)

	if watcher.ID != 0 {

		bytePassword := []byte(form.Password)
		byteHashedPassword := []byte(watcher.Password)
		err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
		checkErr(err, "Invalid password")

		return watcher, nil
	}

	return watcher, errors.New("Create an account")
}

//Signup ...
func (m WatcherModel) Signup(form forms.SignupForm) (watcher Watcher, err error) {
	getDb := db.GetDB()
	getDb.Where(&Watcher{Email: form.Email}).First(&watcher)

	if getDb.NewRecord(watcher) {
		bytePassword := []byte(form.Password)
		hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
		checkErr(err, "Pass hash failed")

		watcher.Email = form.Email
		watcher.Name = form.Name
		watcher.Password = string(hashedPassword)
		getDb.Create(&watcher)

		return watcher, nil
	}

	return watcher, errors.New("Watcher exists")
}

//One ...
func (m WatcherModel) One(watcherID uint) (watcher Watcher) {
	db.GetDB().First(&watcher, watcherID)
	return watcher
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
