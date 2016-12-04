package controllers

import (
	"github.com/vodstv/core/dao"
	"github.com/vodstv/core/models"

	"github.com/gin-gonic/gin"
)

//VodController ...
type VodController struct{}

var vodDao = new(dao.VodDao)

//Find ...
func (ctrl VodController) Find(c *gin.Context) {
	if !checkLogin(c) {
		return
	}

	vodSearch := make(map[string]interface{})
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
	if !checkLogin(c) {
		return
	}

	vta := struct {
		VodID uint `json:"vodId"`
		TagID uint `json:"tagId"`
	}{}
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
	if !checkLogin(c) {
		return
	}

	vod, err := vodDao.Get(getIDParam(c))
	if checkErr(c, err, "Vod get failed") {
		c.JSON(200, vod)
	}
}

//Save ...
func (ctrl VodController) Save(c *gin.Context) {
	if !checkLogin(c) {
		return
	}

	vod := models.Vod{}
	err := c.BindJSON(&vod)
	if checkErr(c, err, "Vod convert failed") {
		err = vodDao.Save(&vod)
		if checkErr(c, err, "Vod save failed") {
			c.JSON(200, vod)
		}
	}
}

//Delete ...
func (ctrl VodController) Delete(c *gin.Context) {
	if !checkLogin(c) {
		return
	}

	vod, err := vodDao.Delete(getIDParam(c))
	if checkErr(c, err, "Vod delete failed") {
		c.JSON(200, vod)
	}
}
