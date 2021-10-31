package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ksrnnb/chat-app-server/adapter/gateway"
	"github.com/ksrnnb/chat-app-server/infrastructure/session"
	"github.com/ksrnnb/chat-app-server/usecase"
)


func getUser(c *gin.Context) {
	interactor := usecase.NewUserInteractor(gateway.NewUserRepository(gateway.NewSqlHandler()))
	res := userController.GetUser(session.NewSession(c), interactor)

	c.JSON(res.Code, res.Params)
}