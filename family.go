package main

import (
	"family/middleware"
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
	p := router.Group("/person")
	p.GET("/get", person.Get)
	p.GET("/list", person.List)
	p.GET("/import", person.Import)

	p.GET("/flush", person.Flush)
	p.GET("/getChildren", person.GetChildren)
	p.GET("/getPosterity", person.GetAllPosterity)

	router.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "index"})
	})
	router.StaticFS("/raoumer", http.Dir("raoumer"))
	router.Run(":8888")
}
