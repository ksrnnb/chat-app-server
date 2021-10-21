package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ksrnnb/chat-app-server/adapter/controller"
	"github.com/ksrnnb/chat-app-server/adapter/gateway"
	"github.com/ksrnnb/chat-app-server/adapter/request"

	"github.com/ksrnnb/chat-app-server/usecase"
)

var loginController controller.LoginController

func init() {
	loginController = controller.NewLoginController()
}

func SetRoute(r *gin.Engine) *gin.Engine {
	r.POST("/login", login)
	return r
}

func login(c *gin.Context) {
	var req request.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	interactor := usecase.NewUserInteractor(gateway.NewUserRepository(gateway.NewSqlHandler()))
	res := loginController.Login(req, interactor)
	c.JSON(res.Code, res.Params)
}