package main

import (
	"family/middleware"
	"family/oauth2"
	"family/person"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.Logger())

	//person相关路由
	p := router.Group("/api/person")
	p.GET("/get", person.Get)
	p.GET("/list", person.List)
	p.GET("/import", person.Import)

	p.GET("/flush", person.Flush)
	p.GET("/getChildren", person.GetChildren)
	p.GET("/getPosterity", person.GetAllPosterity)

	//page info相关路由
	info := router.Group("/info")
	info.StaticFile("/list", "./templates/info/list.html")
	info.LoafHTMLGlob("templates/*")

	//oauth2相关路由
	o := router.Group("/oauth2")
	o.GET("/showcode", oauth2.ShowCode)
	router.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "welcome to sunshijiapu.com.cn"})
	})
	router.StaticFS("/raoumer", http.Dir("raoumer"))
	router.StaticFS("/static", http.Dir("static"))
	router.StaticFile("/dashboard", "./templates/dashboard.html")

	router.Run(":8888")
}
