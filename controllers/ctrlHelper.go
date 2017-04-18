package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func checkErr(c *gin.Context, err error, msg string) bool {
	if err != nil {
		c.JSON(406, gin.H{"Message": msg, "error": err.Error()})
		c.Abort()
		return false
	}
	return true
}

func bindJSONToForm(c *gin.Context, form interface{}) bool {

	if c.BindJSON(&form) != nil {
		c.JSON(406, gin.H{"message": "Invalid json", "form": form})
		c.Abort()
		return false
	}
	return true
}

func getIDParam(c *gin.Context) uint {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	checkErr(c, err, "Not found")
	return uint(id)
}
