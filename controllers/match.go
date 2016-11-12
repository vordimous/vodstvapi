package controllers

import (
	"esvodsApi/forms"
	"esvodsCore/dao"
	"esvodsCore/models"

	"github.com/gin-gonic/gin"
)

//MatchController ...
type MatchController struct{}

var matchDao = new(dao.MatchDao)

//Find ...
func (ctrl MatchController) Find(c *gin.Context) {
	checkLogin(c)

	var matchSearch forms.MatchSearch
	if !bindJSONToForm(c, &matchSearch) {
		return
	}

	match, err := matchDao.Find(matchSearch)
	if checkErr(c, err, "Could not find matchs") {
		c.JSON(200, match)
	}
}

//Get ...
func (ctrl MatchController) Get(c *gin.Context) {
	checkLogin(c)

	match, err := matchDao.Get(getIDParam(c))
	if checkErr(c, err, "Match get failed") {
		c.JSON(200, match)
	}
}

//Save ...
func (ctrl MatchController) Save(c *gin.Context) {
	checkLogin(c)

	var matchForm forms.MatchForm
	if !bindJSONToForm(c, &matchForm) {
		return
	}

	var match = models.Match{}
	var err error
	if matchForm.ID != 0 {
		match, err = matchDao.Get(matchForm.ID)
	}

	err = matchForm.ToModel(&match)
	checkErr(c, err, "Match convert failed")

	err = matchDao.Save(&match)
	if checkErr(c, err, "Match create failed") {
		c.JSON(200, match)
	}
}

//AscVod ...
func (ctrl MatchController) AscVod(c *gin.Context) {
	checkLogin(c)

	var vta forms.MatchVodAsc
	if !bindJSONToForm(c, &vta) {
		return
	}

	if vta.MatchID != 0 && vta.VodID != 0 {
		var match models.Match
		var vod models.Vod
		var err error
		match, err = matchDao.Get(vta.MatchID)
		vod, err = vodDao.Get(vta.VodID)
		if checkErr(c, err, "Could not find items") {
			match.Vods = append(match.Vods, vod)
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

//Delete ...
func (ctrl MatchController) Delete(c *gin.Context) {
	checkLogin(c)

	err := matchDao.Delete(getIDParam(c))
	if checkErr(c, err, "Match delete failed") {
		c.JSON(200, gin.H{"message": "Match deleted"})
	}
}
