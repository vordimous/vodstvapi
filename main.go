package main

import (
	"fmt"

	"esvodsApi/controllers"
	"esvodsCore/dao"
	"esvodsCore/sess"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/ekyoung/gin-nice-recovery"
)

//CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func recoveryHandler(c *gin.Context, err interface{}) {
	c.JSON(500, gin.H{
		"title": "Error",
		"err":   err,
	})
}

func httpTest(c *gin.Context) {
	status := make(map[string]string)
	err := dao.GetDB().DB().Ping()
	if err == nil {
		status["dbConn"] = "true"
	} else {
		status["dbConn"] = err.Error()
	}
	c.JSON(200, status)
}

func main() {
	r := gin.Default()

	//Middlewear
	r.Use(nice.Recovery(recoveryHandler))
	r.Use(CORSMiddleware())
	r.Use(sessions.Sessions("esvods-session", sess.Init()))

	dao.Init()
	dao.DbMigration()

	v1 := r.Group("/v1")
	{
		v1.GET("/status", httpTest)

		/*** START Watcher ***/
		watcher := new(controllers.WatcherController)

		v1.POST("/watcher/login", watcher.Signin)
		v1.POST("/watcher/register", watcher.Signup)
		v1.GET("/watcher/signout", watcher.Signout)
		v1.GET("/me", watcher.Me)

		/*** START Vod ***/
		vod := new(controllers.VodController)

		v1.POST("/vod", vod.Save)
		v1.POST("/vods", vod.Find)
		v1.GET("/vod/:id", vod.Get)
		v1.DELETE("/vod/:id", vod.Delete)
	}

	r.Run(":9000")
}
