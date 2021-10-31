package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ksrnnb/chat-app-server/infrastructure/session"
)

func getToken(c *gin.Context) {
	s := session.NewSession(c)
	res := tokenController.GetToken(s)

	c.JSON(res.Code, res.Params)
}