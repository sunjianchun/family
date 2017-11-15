package oauth2

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowCode(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	code := c.DefaultQuery("code", "")
	state := c.DefaultQuery("state", "")
	access_token := c.DefaultQuery("access_token", "")
	refresh_token := c.DefaultQuery("refresh_token", "")
	expires_in := c.DefaultQuery("expires_in", "")
	token_type := c.DefaultQuery("token_type", "")

	dict := make(map[string]string, 5)
	if id != "" {
		dict["id"] = id
	}
	if state != "" {
		dict["state"] = state
	}
	if access_token != "" {
		dict["access_token"] = access_token
	}
	if refresh_token != "" {
		dict["refresh_token"] = refresh_token
	}
	if code != "" {
		dict["code"] = code
	}
	if expires_in != "" {
		dict["expires_in"] = expires_in
	}
	if token_type != "" {
		dict["token_type"] = token_type
	}
	c.JSON(http.StatusOK, dict)
}
