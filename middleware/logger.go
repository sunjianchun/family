package middleware

import (
	"family/conf"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {

	accesslog := path.Join(conf.BC.Data["base"]["logdir"], "access.log")
	file, _ := os.Create(accesslog)

	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(file)

	if conf.BC.Data["base"]["logdir"] != "product" {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
}

//Logger 输入到日志的文件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		timeStartStr := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
		httpMethod := c.Request.Method
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		contentType := c.ContentType()

		//执行
		c.Next()
		c.Header("Content-Type", contentType)
		if raw != "" {
			path = path + "?" + raw
		}
		latency := time.Since(t)
		code := c.Writer.Status()
		clientIP := c.ClientIP()
		contextLogger := log.WithFields(log.Fields{
			"time":     timeStartStr,
			"method":   httpMethod,
			"code":     code,
			"path":     path,
			"cost":     latency,
			"clientip": clientIP,
		})

		if code == 200 || code == 301 || code == 302 {
			contextLogger.Info()
		} else if code == 404 {
			//if strings.Contains(contentType, "text/html") {
			c.Redirect(http.StatusMovedPermanently, "/notfound")
			//}
			contextLogger.Warn()
			return
		} else if code == 403 || code == 500 || code == 502 {
			contextLogger.Error()
		} else {
			contextLogger.Info()
		}
	}
}

//Login 登录
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("userID")
		if userID == nil && c.Request.URL.Path != "/login" {
			c.Redirect(http.StatusMovedPermanently, "/login")
		}

	}
}
