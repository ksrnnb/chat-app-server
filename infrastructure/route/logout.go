package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ksrnnb/chat-app-server/infrastructure/session"
)

func logout(c *gin.Context) {
	s := session.NewSession(c)
	_, ok := s.Get("userId").(int)

	if !ok {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	res := logoutController.Logout(s)

	c.JSON(res.Code, res.Params)
}