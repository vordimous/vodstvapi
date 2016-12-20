package controllers

import (
	"vodstv/core/dao"
	"vodstv/core/models"

	"github.com/gin-gonic/gin"
)

//FeedController ...
type FeedController struct{}

var feedDao = new(dao.FeedDao)

//Find ...
func (ctrl FeedController) Find(c *gin.Context) {
	feedSearch := make(map[string]interface{})
	if !bindJSONToForm(c, &feedSearch) {
		return
	}

	feed, err := feedDao.Find(feedSearch)
	if checkErr(c, err, "Could not find feeds") {
		c.JSON(200, feed)
	}
}

//Get ...
func (ctrl FeedController) Get(c *gin.Context) {
	feed, err := feedDao.Get(getIDParam(c))
	if checkErr(c, err, "Feed get failed") {
		c.JSON(200, feed)
	}
}

//Save ...
func (ctrl FeedController) Save(c *gin.Context) {
	feed := models.Feed{}
	err := c.BindJSON(&feed)
	if checkErr(c, err, "Feed convert failed") {
		err = feedDao.Save(&feed)
		if checkErr(c, err, "Feed save failed") {
			c.JSON(200, feed)
		}
	}
}

//Delete ...
func (ctrl FeedController) Delete(c *gin.Context) {
	err := feedDao.Delete(getIDParam(c))
	if checkErr(c, err, "Feed delete failed") {
		c.JSON(200, gin.H{"message": "Feed deleted"})
	}
}
