package main

import (
	"fmt"
	"os"

	"esvodsApi/controllers"
	"esvodsApi/policies"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/ekyoung/gin-nice-recovery"

	"github.com/vodstv/core/dao"
	"github.com/vodstv/core/sess"
)

//CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if origin := c.Request.Header.Get("Origin"); origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}
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

	v1 := r.Group("/v1")
	{
		v1.GET("/status", httpTest)

		/*** START Watcher ***/
		watcher := new(controllers.WatcherController)

		v1.POST("/login", watcher.Signin)
		v1.POST("/register", watcher.Signup)
		v1.GET("/signout", watcher.Signout)
		v1.GET("/me", watcher.Me)

		v1.POST("/watcher", policies.ReqAuth(), policies.SelfOrAdmin(), watcher.Save)
		v1.POST("/watchers", policies.ReqAuth(), policies.Admin(), watcher.Find)
		v1.GET("/watcher/:id", policies.ReqAuth(), policies.SelfOrAdmin(), watcher.Get)
		v1.DELETE("/watcher/:id", policies.ReqAuth(), policies.SelfOrAdmin(), watcher.Delete)

		/*** START Vod ***/
		vod := new(controllers.VodController)

		v1.POST("/vod", vod.Save)
		v1.POST("/vodAscTag", vod.AscTag)
		v1.POST("/vods", vod.Find)
		v1.GET("/vod/:id", vod.Get)
		v1.DELETE("/vod/:id", vod.Delete)

		/*** START Match ***/
		match := new(controllers.MatchController)

		v1.POST("/match", match.Save)
		v1.POST("/matchAscVod", match.AscVod)
		v1.POST("/matches", match.Find)
		v1.GET("/match/:id", match.Get)
		v1.DELETE("/match/:id", match.Delete)

		/*** START Tag ***/
		tag := new(controllers.TagController)

		v1.POST("/tag", tag.Save)
		v1.POST("/tags", tag.Find)
		v1.GET("/tag/:id", tag.Get)
		v1.DELETE("/tag/:id", tag.Delete)

		/*** START Feed ***/
		feed := new(controllers.FeedController)

		v1.POST("/feed", policies.ReqAuth(), feed.Save)
		v1.POST("/feeds", policies.ReqAuth(), feed.Find)
		v1.GET("/feed/:id", policies.ReqAuth(), feed.Get)
		v1.DELETE("/feed/:id", policies.ReqAuth(), feed.Delete)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "9000"
	}

	r.Run(":" + port)
}
