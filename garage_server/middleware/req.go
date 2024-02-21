package middleware

import (
	"github.com/gin-gonic/gin"
)

func (m *middleware) RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// savedCtx := c.Request.Context()

		var key = []interface{}{
			"http_method", c.Request.Method,
			"path", c.Request.RequestURI,
			"client_ip", c.ClientIP(),
			"status", c.Writer.Status(),
		}

		c.Next()
		m.logger.Infow("", key...)
	}

}
