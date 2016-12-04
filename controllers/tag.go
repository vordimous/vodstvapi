package controllers

import (
	"github.com/vodstv/core/dao"
	"github.com/vodstv/core/models"

	"github.com/gin-gonic/gin"
)

//TagController ...
type TagController struct{}

var tagDao = new(dao.TagDao)

//Find ...
func (ctrl TagController) Find(c *gin.Context) {
	if !checkLogin(c) {
		return
	}

	tagSearch := make(map[string]interface{})
	if !bindJSONToForm(c, &tagSearch) {
		return
	}

	tag, err := tagDao.Find(tagSearch)
	if checkErr(c, err, "Could not find tags") {
		c.JSON(200, tag)
	}
}

//Get ...
func (ctrl TagController) Get(c *gin.Context) {
	if !checkLogin(c) {
		return
	}

	tag, err := tagDao.Get(getIDParam(c))
	if checkErr(c, err, "Tag get failed") {
		c.JSON(200, tag)
	}
}

//Save ...
func (ctrl TagController) Save(c *gin.Context) {
	if !checkLogin(c) {
		return
	}

	tag := models.Tag{}
	err := c.BindJSON(&tag)
	if checkErr(c, err, "Tag convert failed") {
		err = tagDao.Save(&tag)
		if checkErr(c, err, "Tag create failed") {
			c.JSON(200, tag)
		}
	}
}

//Delete ...
func (ctrl TagController) Delete(c *gin.Context) {
	if !checkLogin(c) {
		return
	}

	tag, err := tagDao.Delete(getIDParam(c))
	if checkErr(c, err, "Tag delete failed") {
		c.JSON(200, tag)
	}
}
