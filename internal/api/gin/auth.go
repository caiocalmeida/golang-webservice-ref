package api

import (
	"github.com/gin-gonic/gin"
)

func requireApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		if apiKey := c.GetHeader("X-API-KEY"); apiKey != "example" {
			c.String(401, "Please provide a valid key in the X-API-KEY request header. (Since this is just an example API, use \"example\" as the key)")
			c.Abort()
			return
		}

		c.Next()
	}
}

func forbidRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(403, "This API endpoint is unavailable in order to prevent abuse.")
		c.Abort()
	}
}
