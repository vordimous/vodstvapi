package policies

import (
	"github.com/vodstv/core/sess"

	"github.com/gin-gonic/gin"
)

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

//Perm ...
func Perm(perm string) gin.HandlerFunc {
	return func(c *gin.Context) {
		notAllowed := true
		if notAllowed {
			c.JSON(405, gin.H{"err": "Not allowed"})
		} else {
			c.Next()
		}
	}
}
