package main

import (
	inner_http "family/http"
	"family/middleware"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	store := sessions.NewCookieStore([]byte("secret"))
	store.Options(sessions.Options{MaxAge: 3600})
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(sessions.Sessions("mySession", store))
	router.Use(middleware.Login())
	router.Use(middleware.Logger())
	router.StaticFS("/static", http.Dir("static"))
	router.LoadHTMLGlob("templates/*/*")

	//person相关路由
	p := router.Group("/api/person")
	p.GET("/get", inner_http.Get)
	p.GET("/list", inner_http.List)
	p.POST("/add", inner_http.Add)
	p.GET("/getUserInfo", inner_http.GetUserInfo)
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
	info.GET("/flexible", inner_http.Infoflexible)

	//dashboard相关路由
	dashboard := router.Group("/dashboard")
	dashboard.GET("/", inner_http.DashBoardIndex)

	//oauth2相关路由
	o := router.Group("/oauth2")
	o.GET("/showcode", inner_http.ShowCode)

	//login
	router.GET("/login", inner_http.Login)
	router.POST("/login", inner_http.Login)
	router.GET("/logout", inner_http.Logout)
	//404
	router.GET("/notfound1", inner_http.NotFound)

	router.Run(":8888")
}
