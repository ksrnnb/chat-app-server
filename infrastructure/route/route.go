package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ksrnnb/chat-app-server/adapter/controller"
	"github.com/ksrnnb/chat-app-server/adapter/gateway"
	"github.com/ksrnnb/chat-app-server/adapter/request"
	"github.com/ksrnnb/chat-app-server/infrastructure/session"
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

	if !res.IsSuccessful() {
		c.JSON(res.Code, res.Params)
		return
	}

	// セッションにuserIdを保存。
	// s.Setを実行すると、sessionIdがSet-Cookieヘッダに付与される
	s := session.NewSession(c)
	s.Set("userId", res.Params["id"])

	if err := s.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(res.Code, res.Params)
}
