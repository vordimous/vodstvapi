package dao

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"esvodsApi/forms"
	"esvodsApi/models"
)

//WatcherDao ...
type WatcherDao struct{}

//Signin ...
func (d WatcherDao) Signin(form forms.SigninForm) (watcher models.Watcher, err error) {
	GetDB().Where(&models.Watcher{Email: form.Email}).First(&watcher)

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
func (d WatcherDao) Signup(form forms.SignupForm) (watcher models.Watcher, err error) {
	GetDB().Where(&models.Watcher{Email: form.Email}).First(&watcher)

	if GetDB().NewRecord(watcher) {
		bytePassword := []byte(form.Password)
		hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
		checkErr(err, "Pass hash failed")

		watcher.Email = form.Email
		watcher.Name = form.Name
		watcher.Password = string(hashedPassword)
		GetDB().Create(&watcher)

		return watcher, nil
	}

	return watcher, errors.New("Watcher exists")
}

//One ...
func (d WatcherDao) One(watcherID uint) (watcher models.Watcher) {
	GetDB().First(&watcher, watcherID)
	return watcher
}
