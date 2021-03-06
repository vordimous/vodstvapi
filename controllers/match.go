package controllers

import (
	"vodstv/core/dao"
	"vodstv/core/models"

	"github.com/gin-gonic/gin"
)

//MatchController ...
type MatchController struct{}

var matchDao = new(dao.MatchDao)

//Find ...
func (ctrl MatchController) Find(c *gin.Context) {
	matchSearch := make(map[string]interface{})
	if !bindJSONToForm(c, &matchSearch) {
		return
	}

	match, err := matchDao.Find(matchSearch)
	if checkErr(c, err, "Could not find matchs") {
		c.JSON(200, match)
	}
}

//Query ...
func (ctrl MatchController) Query(c *gin.Context) {
	matchQuery := struct {
		TagIDs []uint `json:"tagIds"`
	}{}

	if !bindJSONToForm(c, &matchQuery) {
		return
	}

	match, err := matchDao.Query(matchQuery.TagIDs)
	if checkErr(c, err, "Could not find matchs") {
		c.JSON(200, match)
	}
}

//AscTag ...
func (ctrl MatchController) AscTag(c *gin.Context) {
	vta := struct {
		MatchID uint `json:"matchId"`
		TagID   uint `json:"tagId"`
	}{}
	if !bindJSONToForm(c, &vta) {
		return
	}

	if vta.MatchID != 0 && vta.TagID != 0 {
		var match models.Match
		var tag models.Tag
		var err error
		match, err = matchDao.Get(vta.MatchID)
		tag, err = tagDao.Get(vta.TagID)
		if checkErr(c, err, "Could not find items") {
			match.Tags = append(match.Tags, tag)
			matchDao.Save(&match)

			if checkErr(c, err, "Could not save match") {
				c.JSON(200, match)
			}
		}
	} else {
		c.JSON(406, gin.H{"Message": "Must supply both IDs", "form": vta})
		c.Abort()
		return
	}

}

//Get ...
func (ctrl MatchController) Get(c *gin.Context) {
	match, err := matchDao.Get(getIDParam(c))
	if checkErr(c, err, "Match get failed") {
		c.JSON(200, match)
	}
}

//Save ...
func (ctrl MatchController) Save(c *gin.Context) {
	match := models.Match{}
	err := c.BindJSON(&match)
	if checkErr(c, err, "Match convert failed") {
		err = matchDao.Save(&match)
		if checkErr(c, err, "Match save failed") {
			c.JSON(200, match)
		}
	}
}

//AscVod ...
func (ctrl MatchController) AscVod(c *gin.Context) {
	mva := struct {
		MatchID uint `json:"matchId"`
		VodID   uint `json:"vodId"`
	}{}
	if !bindJSONToForm(c, &mva) {
		return
	}

	if mva.MatchID != 0 && mva.VodID != 0 {
		var match models.Match
		var vod models.Vod
		var err error
		match, err = matchDao.Get(mva.MatchID)
		vod, err = vodDao.Get(mva.VodID)
		if checkErr(c, err, "Could not find items") {
			match.Vods = append(match.Vods, vod)
			matchDao.Save(&match)

			if checkErr(c, err, "Could not save match") {
				c.JSON(200, match)
			}
		}
	} else {
		c.JSON(406, gin.H{"Message": "Must supply both IDs", "form": mva})
		c.Abort()
		return
	}

}

//Delete ...
func (ctrl MatchController) Delete(c *gin.Context) {
	err := matchDao.Delete(getIDParam(c))
	if checkErr(c, err, "Match delete failed") {
		c.JSON(200, gin.H{"message": "Match deleted", "success": true})
	}
}
