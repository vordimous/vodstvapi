package controllers

import (
	"esvodsApi/forms"
	"esvodsCore/dao"
	"esvodsCore/models"

	"github.com/gin-gonic/gin"
)

//VodController ...
type VodController struct{}

var vodDao = new(dao.VodDao)

//Find ...
func (ctrl VodController) Find(c *gin.Context) {
	checkLogin(c)

	var vodSearch forms.VodSearch
	bindJSONToForm(c, &vodSearch)

	data, err := vodDao.Find(vodSearch)
	checkErr(c, err, "Could not find vods")

	c.JSON(200, data)
}

//Get ...
func (ctrl VodController) Get(c *gin.Context) {
	checkLogin(c)

	vod, err := vodDao.Get(getIDParam(c))
	checkErr(c, err, "Vod get failed")

	c.JSON(200, vod)
}

//Save ...
func (ctrl VodController) Save(c *gin.Context) {
	checkLogin(c)

	var vodForm forms.VodForm
	bindJSONToForm(c, &vodForm)

	var vod = models.Vod{} 
	var err error
	if vodForm.ID != 0 {
		vod, err = vodDao.Get(vodForm.ID)
	}

	err = vodForm.ToModel(&vod)
	checkErr(c, err, "Vod convert failed")

	err = vodDao.Save(&vod)
	checkErr(c, err, "Vod create failed")

	c.JSON(200, vod)
}

//Delete ...
func (ctrl VodController) Delete(c *gin.Context) {
	checkLogin(c)

	err := vodDao.Delete(getIDParam(c))
	checkErr(c, err, "Vod delete failed")

	c.JSON(200, gin.H{"message": "Vod deleted"})
}
