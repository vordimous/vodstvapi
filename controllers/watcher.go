package controllers

import (
	"fmt"

	"esvodsApi/forms"
	"esvodsApi/models"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

//WatcherController ...
type WatcherController struct{}

var watcherModel = new(models.WatcherModel)

//getWatcherID ...
func getWatcherID(c *gin.Context) uint {
	session := sessions.Default(c)
	watcherID := session.Get("watcher_id")
	if watcherID != nil {
		return models.ConvertToUInt(watcherID)
	}
	return 0
}

//getSessionWatcherInfo ...
func getSessionWatcherInfo(c *gin.Context) (watcherSessionInfo models.WatcherSessionInfo) {
	session := sessions.Default(c)
	watcherID := session.Get("watcher_id")
	if watcherID != nil {
		watcherSessionInfo.ID = models.ConvertToUInt(watcherID)
		watcherSessionInfo.Name = session.Get("watcher_name").(string)
		watcherSessionInfo.Email = session.Get("watcher_email").(string)
	}
	return watcherSessionInfo
}

//Signin ...
func (ctrl WatcherController) Signin(c *gin.Context) {
	var signinForm forms.SigninForm

	if c.BindJSON(&signinForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": signinForm})
		c.Abort()
		return
	}

	watcher, err := watcherModel.Signin(signinForm)
	if err == nil {
		session := sessions.Default(c)
		session.Set("watcher_id", watcher.ID)
		session.Set("watcher_email", watcher.Email)
		session.Set("watcher_name", watcher.Name)
		session.Save()

		c.JSON(200, getSessionWatcherInfo(c))
	} else {
		c.JSON(406, gin.H{"message": "Invalid signin details", "error": err.Error()})
	}

}

//Signup ...
func (ctrl WatcherController) Signup(c *gin.Context) {
	var signupForm forms.SignupForm

	if c.BindJSON(&signupForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": signupForm})
		c.Abort()
		return
	}

	watcher, err := watcherModel.Signup(signupForm)

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
		watcherID := getWatcherID(c)
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
	c.JSON(200, getSessionWatcherInfo(c))
}
