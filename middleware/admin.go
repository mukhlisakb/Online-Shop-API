package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: get header auth
		key := os.Getenv("ADMIN_SECRET")
		// TODO: validate header with admin password
		auth := c.Request.Header.Get("Authorization")
		if auth == " " {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		if auth != key {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
		// TODO: continue process to handler
	}
}
