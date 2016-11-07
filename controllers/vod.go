package controllers

import (
	"esvodsApi/dao"
	"esvodsApi/forms"

	"github.com/gin-gonic/gin"
)

//VodController ...
type VodController struct{}

var vodDao = new(dao.VodDao)

//Create ...
func (ctrl VodController) Create(c *gin.Context) {
	checkLogin(c)

	var vodForm forms.VodForm
	bindJsonToForm(c, &vodForm)

	vodID, err := vodDao.Create(vodForm)

	checkErr(c, err, "Vod create failed")

	c.JSON(200, gin.H{"message": "Vod created", "id": vodID})
}

//All ...
func (ctrl VodController) All(c *gin.Context) {
	checkLogin(c)

	data, err := vodDao.All()

	checkErr(c, err, "Could not get the vods")

	c.JSON(200, data)
}

//One ...
func (ctrl VodController) One(c *gin.Context) {
	checkLogin(c)

	id := getIDParam(c)

	var vodForm forms.VodForm
	bindJsonToForm(c, &vodForm)

	data, err := vodDao.One(id)
	checkErr(c, err, "Vod get failed")

	c.JSON(200, data)
}

//Update ...
func (ctrl VodController) Update(c *gin.Context) {
	checkLogin(c)
	var vodForm forms.VodForm
	bindJsonToForm(c, &vodForm)

	err := vodDao.Update(vodForm)
	checkErr(c, err, "Vod update failed")

	c.JSON(200, gin.H{"message": "Vod updated"})
}

//Delete ...
func (ctrl VodController) Delete(c *gin.Context) {
	checkLogin(c)

	id := getIDParam(c)

	var vodForm forms.VodForm
	bindJsonToForm(c, &vodForm)

	err := vodDao.Delete(id)
	checkErr(c, err, "Vod delete failed")

	c.JSON(200, gin.H{"message": "Vod deleted"})
}
