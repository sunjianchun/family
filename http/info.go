package http

import (
	"net/http"

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
