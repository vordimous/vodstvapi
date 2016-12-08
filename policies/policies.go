package policies

import (
	"strconv"

	"github.com/vodstv/core/sess"

	"github.com/gin-gonic/gin"
)

func getIDParam(c *gin.Context) uint {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	return uint(id)
}

//ReqAuth ...
func ReqAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		watcherID := sess.GetWatcherID(c)
		if watcherID == 0 {
			c.JSON(403, gin.H{"err": "Please login first"})
			c.Abort()
		} else {
			c.Next()
		}
	}
}

//Admin ...
func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := sess.GetSessionWatcherInfo(c)
		if w.IsAdmin {
			c.Next()
		} else {
			c.JSON(405, gin.H{"err": "Must be admin"})
			c.Abort()
		}
	}
}

//SelfOrAdmin ...
func SelfOrAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := sess.GetSessionWatcherInfo(c)
		rID := getIDParam(c)
		if (w.ID != rID) || w.IsAdmin {
			c.Next()
		} else {
			c.JSON(405, gin.H{"err": "Must be owner or admin"})
			c.Abort()
		}
	}
}

//Perm ...
func Perm(perm string) gin.HandlerFunc {
	return func(c *gin.Context) {
		notAllowed := true
		if notAllowed {
			c.JSON(405, gin.H{"err": "Must be " + perm})
		} else {
			c.Next()
		}
	}
}
