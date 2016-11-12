package controllers

import (
	"esvodsApi/forms"
	"esvodsCore/dao"
	"esvodsCore/models"

	"github.com/gin-gonic/gin"
)

//TagController ...
type TagController struct{}

var tagDao = new(dao.TagDao)

//Find ...
func (ctrl TagController) Find(c *gin.Context) {
	checkLogin(c)

	var tagSearch forms.TagSearch
	if !bindJSONToForm(c, &tagSearch){
		return
	}

	tag, err := tagDao.Find(tagSearch)
	if checkErr(c, err, "Could not find tags") {
		c.JSON(200, tag)
	}
}

//Get ...
func (ctrl TagController) Get(c *gin.Context) {
	checkLogin(c)

	tag, err := tagDao.Get(getIDParam(c))
	if checkErr(c, err, "Tag get failed") {
		c.JSON(200, tag)
	}
}

//Save ...
func (ctrl TagController) Save(c *gin.Context) {
	checkLogin(c)

	var tagForm forms.TagForm
	if !bindJSONToForm(c, &tagForm){
		return
	}

	var tag = models.Tag{}
	var err error
	if tagForm.ID != 0 {
		tag, err = tagDao.Get(tagForm.ID)
	}

	err = tagForm.ToModel(&tag)
	checkErr(c, err, "Tag convert failed")

	err = tagDao.Save(&tag)
	if checkErr(c, err, "Tag create failed") {
		c.JSON(200, tag)
	}
}

//Delete ...
func (ctrl TagController) Delete(c *gin.Context) {
	checkLogin(c)

	err := tagDao.Delete(getIDParam(c))
	if checkErr(c, err, "Tag delete failed") {
		c.JSON(200, gin.H{"message": "Tag deleted"})
	}
}
