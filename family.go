package main

import (
	inner_http "family/http"
	"family/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.Logger())
	router.LoadHTMLGlob("templates/*/*")

	//person相关路由
	p := router.Group("/api/person")
	p.GET("/get", inner_http.Get)
	p.GET("/list", inner_http.List)
	p.GET("/import", inner_http.Import)

	p.GET("/flush", inner_http.Flush)
	p.GET("/getChildren", inner_http.GetChildren)
	p.GET("/getPosterity", inner_http.GetAllPosterity)

	//page info相关路由
	info := router.Group("/info")
	info.GET("/get", inner_http.PersonDetail)
	info.GET("/list", inner_http.PersonList)
	info.GET("/add", inner_http.PersonAdd)
	info.GET("/personalist", inner_http.PersonalList)
	info.GET("/tree", inner_http.InfoTree)

	info.StaticFile("/dashboard", "./templates/dashboard.html")

	//oauth2相关路由
	o := router.Group("/oauth2")
	o.GET("/showcode", inner_http.ShowCode)
	router.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "welcome to sunshijiapu.com.cn"})
	})
	router.StaticFS("/static", http.Dir("static"))

	router.Run(":8888")
}
