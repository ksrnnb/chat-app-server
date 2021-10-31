package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ksrnnb/chat-app-server/infrastructure/session"
)

type csrfRequest struct {
	csrfToken string
}

func Csrf() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := session.NewSession(c)
		token, ok := s.Get("csrf_token").(string)

		if !ok {
			abort(c)
			return
		}

		var req csrfRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			abort(c)
			return
		}

		if token != req.csrfToken {
			abort(c)
			return
		}

		c.Next()
	}
}

func abort(c *gin.Context) {
	c.JSON(http.StatusUnprocessableEntity, nil)
	c.Abort()
}
