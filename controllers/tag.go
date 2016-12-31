package controllers

import (
	"vodstv/core/dao"
	"vodstv/core/models"

	"github.com/gin-gonic/gin"
)

//TagController ...
type TagController struct{}

var tagDao = new(dao.TagDao)

//FindByTags ...
func (ctrl TagController) FindByTags(c *gin.Context) {
	tagSearch := struct {
		TagIDs []uint `json:"tagIds"`
	}{}
	if !bindJSONToForm(c, &tagSearch) {
		return
	}

	tag, err := tagDao.FindByTags(tagSearch.TagIDs)
	if checkErr(c, err, "Could not find tags") {
		c.JSON(200, tag)
	}
}

//Find ...
func (ctrl TagController) Find(c *gin.Context) {
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
	tag, err := tagDao.Get(getIDParam(c))
	if checkErr(c, err, "Tag get failed") {
		c.JSON(200, tag)
	}
}

//Save ...
func (ctrl TagController) Save(c *gin.Context) {
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
	tag, err := tagDao.Delete(getIDParam(c))
	if checkErr(c, err, "Tag delete failed") {
		c.JSON(200, tag)
	}
}
