package controllers

import (
	"esvodsApi/dao"
	"strconv"

	"github.com/gin-gonic/gin"
)

func checkErr(c *gin.Context, err error, msg string) {
	if err != nil {
		c.JSON(406, gin.H{"Message": msg, "error": err.Error()})
		c.Abort()
		return
	}
}

func checkLogin(c *gin.Context) {
	watcherID := dao.GetWatcherID(c)

	if watcherID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}
}

func bindJsonToForm(c *gin.Context, form interface{}) {

	if c.BindJSON(&form) != nil {
		c.JSON(406, gin.H{"message": "Invalid json", "form": form})
		c.Abort()
		return
	}
}

func getIDParam(c *gin.Context) uint {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err == nil {
		return uint(id)
	}

	checkErr(c, err, "Not found")
	return 0
}
