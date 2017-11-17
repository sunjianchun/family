package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DashBoardIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", gin.H{})
}
