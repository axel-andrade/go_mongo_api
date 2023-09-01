package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Cache(duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Cache-Control", "max-age="+duration.String())
		c.Next()
	}
}
