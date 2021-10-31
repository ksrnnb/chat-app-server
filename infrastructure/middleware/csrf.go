package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/ksrnnb/chat-app-server/infrastructure/session"
)

type csrfRequest struct {
	CsrfToken string `json:"token"`
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
		if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
			abort(c)
			return
		}

		if token != req.CsrfToken {
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
