package http

import (
	"family/conf"
	"family/db"
	"family/util"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func PersonDetail(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "id 不能为空"})
	} else {
		c.HTML(http.StatusOK, "detail.html", gin.H{"userID": id})
	}
}

func PersonAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "add.html", gin.H{})
}

func PersonList(c *gin.Context) {
	c.HTML(http.StatusOK, "list.html", gin.H{})
}

func PersonalList(c *gin.Context) {
	c.HTML(http.StatusOK, "personalist.html", gin.H{})
}

func InfoTree(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "id 不能为空"})
	} else {
		c.HTML(http.StatusOK, "tree.html", gin.H{"userID": id})
	}
}

func Infoflexible(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "id 不能为空"})
	} else {
		c.HTML(http.StatusOK, "book.html", gin.H{"userID": id})
	}
}

func Login(c *gin.Context) {
	var person conf.Person
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	} else {
		err := c.ShouldBind(&person)
		if err != nil {
			c.HTML(http.StatusOK, "login.html", gin.H{"msg": "错误的表单提交"})
		} else {
			sql := "select id from person where name=? and password=?;"
			newdb := db.NewDB(sql)
			result := newdb.Do(conf.Query, person.Name, util.Md5Encode(person.Password))
			if len(result) > 0 {
				sesson := sessions.Default(c)
				sesson.Set("userID", result[0]["id"].(string))
				sesson.Save()
				c.Redirect(301, "/info/get?id="+result[0]["id"].(string))
			} else {
				c.HTML(http.StatusOK, "login.html", gin.H{"msg": "用户名密码错误"})
			}
		}
	}
}

func NotFound(c *gin.Context) {
	c.HTML(http.StatusOK, "404.html", gin.H{})
}

func Logout(c *gin.Context) {
	sesson := sessions.Default(c)
	sesson.Delete("userID")
	c.Redirect(http.StatusPermanentRedirect, "/login")
}
