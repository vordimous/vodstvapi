package controllers

import (
	"fmt"

	"github.com/vodstv/core/dao"
	"github.com/vodstv/core/sess"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

//WatcherController ...
type WatcherController struct{}

var watcherDao = new(dao.WatcherDao)

//Signin ...
func (ctrl WatcherController) Signin(c *gin.Context) {
	signinForm := dao.SigninForm{}

	if c.BindJSON(&signinForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": signinForm})
		c.Abort()
		return
	}

	watcher, err := watcherDao.Signin(signinForm)
	if err == nil {
		session := sessions.Default(c)
		session.Set("watcher_id", watcher.ID)
		session.Set("watcher_email", watcher.Email)
		session.Set("watcher_name", watcher.Name)
		session.Save()

		c.JSON(200, sess.GetSessionWatcherInfo(c))
	} else {
		c.JSON(406, gin.H{"message": "Invalid signin details", "error": err.Error()})
	}

}

//Signup ...
func (ctrl WatcherController) Signup(c *gin.Context) {
	signupForm := dao.SignupForm{}

	if c.BindJSON(&signupForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": signupForm})
		c.Abort()
		return
	}

	watcher, err := watcherDao.Signup(signupForm)

	if err != nil {
		c.JSON(406, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	if watcher.ID > 0 {
		session := sessions.Default(c)
		session.Set("watcher_id", watcher.ID)
		session.Set("watcher_email", watcher.Email)
		session.Set("watcher_name", watcher.Name)
		session.Save()
		c.JSON(200, watcher)
	} else {
		c.JSON(406, gin.H{"message": "Could not signup this watcher", "error": err.Error()})
	}

}

//Signout ...
func (ctrl WatcherController) Signout(c *gin.Context) {
	session := sessions.Default(c)
	if session != nil {
		watcherID := sess.GetWatcherID(c)
		session.Clear()
		session.Save()
		fmt.Println("Logged out:", watcherID)
		c.JSON(200, gin.H{"message": "Signed out..."})
	} else {
		c.JSON(200, gin.H{"message": "Already signed out..."})
	}
}

//Me ...
func (ctrl WatcherController) Me(c *gin.Context) {
	c.JSON(200, sess.GetSessionWatcherInfo(c))
}
