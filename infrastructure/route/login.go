package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ksrnnb/chat-app-server/adapter/gateway"
	"github.com/ksrnnb/chat-app-server/adapter/request"
	"github.com/ksrnnb/chat-app-server/infrastructure/session"
	"github.com/ksrnnb/chat-app-server/usecase"
)

func login(c *gin.Context) {
	var req request.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: コンテナ管理
	sqlHandler := gateway.NewSqlHandler()
	defer sqlHandler.Close()
	interactor := usecase.NewUserInteractor(gateway.NewUserRepository(sqlHandler))
	res := loginController.Login(req, interactor)

	if !res.IsSuccessful() {
		c.JSON(res.Code, res.Params)
		return
	}

	params, ok := res.Params.(map[string]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	// セッションにuserIdを保存。
	// s.Setを実行すると、sessionIdがSet-Cookieヘッダに付与される
	s := session.NewSession(c)
	s.Set("userId", params["id"])

	if err := s.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(res.Code, res.Params)
}