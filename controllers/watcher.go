package controllers

import (
	"fmt"

	"vodstv/core/dao"
	"vodstv/core/models"
	"vodstv/core/sess"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

//WatcherController ...
type WatcherController struct{}

var watcherDao = new(dao.WatcherDao)

//Auth

func setSession(c *gin.Context, w models.Watcher) {
	session := sessions.Default(c)
	session.Set("watcher_id", w.ID)
	session.Set("watcher_email", w.Email)
	session.Set("watcher_username", w.Username)
	session.Set("watcher_is_admin", "true")
	session.Save()
}

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
		setSession(c, watcher)

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
		setSession(c, watcher)
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

//Actions
//AddFeed ...
func (ctrl WatcherController) AddFeed(c *gin.Context) {
	feed := models.Feed{}
	err := c.BindJSON(&feed)
	if checkErr(c, err, "Feed convert failed") {
		err = feedDao.Save(&feed)
		WID := getIDParam(c)

		if checkErr(c, err, "Could not save feed") && feed.ID != 0 && WID != 0 {
			var watcher models.Watcher
			var err error
			watcher, err = watcherDao.Get(WID)

			if checkErr(c, err, "Could not get watcher") {
				watcher.Feeds = append(watcher.Feeds, feed)
				watcherDao.Save(&watcher)

				if checkErr(c, err, "Could not save watcher") {
					c.JSON(200, watcher)
					return
				}
			}
		}
	}
}

//Crud
//Find ...
func (ctrl WatcherController) Find(c *gin.Context) {
	watcherSearch := make(map[string]interface{})
	if !bindJSONToForm(c, &watcherSearch) {
		return
	}

	watcher, err := watcherDao.Find(watcherSearch)
	if checkErr(c, err, "Could not find watchers") {
		c.JSON(200, watcher)
	}
}

//Get ...
func (ctrl WatcherController) Get(c *gin.Context) {
	watcher, err := watcherDao.Get(getIDParam(c))
	if checkErr(c, err, "Watcher get failed") {
		c.JSON(200, watcher)
	}
}

//Save ...
func (ctrl WatcherController) Save(c *gin.Context) {
	watcher := models.Watcher{}
	err := c.BindJSON(&watcher)
	if checkErr(c, err, "Watcher convert failed") {
		err = watcherDao.Save(&watcher)
		if checkErr(c, err, "Watcher save failed") {
			c.JSON(200, watcher)
		}
	}
}

//Delete ...
func (ctrl WatcherController) Delete(c *gin.Context) {
	err := watcherDao.Delete(getIDParam(c))
	if checkErr(c, err, "Watcher delete failed") {
		c.JSON(200, gin.H{"message": "Watcher deleted"})
	}
}
