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
	if !bindJSONToForm(c, &vodSearch) {
		return
	}

	vod, err := vodDao.Find(vodSearch)
	if checkErr(c, err, "Could not find vods") {
		c.JSON(200, vod)
	}
}

//AscTag ...
func (ctrl VodController) AscTag(c *gin.Context) {
	checkLogin(c)

	var vta forms.VodTagAsc
	if !bindJSONToForm(c, &vta) {
		return
	}

	if vta.VodID != 0 && vta.TagID != 0 {
		var vod models.Vod
		var tag models.Tag
		var err error
		vod, err = vodDao.Get(vta.VodID)
		tag, err = tagDao.Get(vta.TagID)
		if checkErr(c, err, "Could not find items") {
			vod.Tags = append(vod.Tags, tag)
			vodDao.Save(&vod)

			if checkErr(c, err, "Could not save vod") {
				c.JSON(200, vod)
			}
		}
	} else {
		c.JSON(406, gin.H{"Message": "Must supply both IDs", "form": vta})
		c.Abort()
		return
	}

}

//Get ...
func (ctrl VodController) Get(c *gin.Context) {
	checkLogin(c)

	vod, err := vodDao.Get(getIDParam(c))
	if checkErr(c, err, "Vod get failed") {
		c.JSON(200, vod)
	}
}

//Save ...
func (ctrl VodController) Save(c *gin.Context) {
	checkLogin(c)

	vodForm := forms.VodForm{}
	if !bindJSONToForm(c, &vodForm) {
		return
	}

	vod := models.Vod{}
	if vodForm.ID != 0 {
		//todo: check for delete

		found, e := vodDao.Get(vodForm.ID)
		if e == nil {
			vod = found
		}
	}

	err := vodForm.ToModel(&vod)
	if !checkErr(c, err, "Vod convert failed") {
		return
	}

	err = vodDao.Save(&vod)
	if checkErr(c, err, "Vod create failed") {
		c.JSON(200, vod)
	}
}

//Delete ...
func (ctrl VodController) Delete(c *gin.Context) {
	checkLogin(c)

	err := vodDao.Delete(getIDParam(c))
	if checkErr(c, err, "Vod delete failed") {
		c.JSON(200, gin.H{"message": "Vod deleted"})
	}
}
