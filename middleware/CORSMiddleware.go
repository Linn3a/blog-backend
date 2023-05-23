package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(c.Request.Header.Get("Origin"))
		origin := c.Request.Header.Get("Origin")
		log.Println(len(origin))

		//pro
		if len(origin) > 23 {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "http://124.220.198.163:81")
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "http://124.220.198.163")
		}

		//dev
		//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
