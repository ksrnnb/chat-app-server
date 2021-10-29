package route

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ksrnnb/chat-app-server/adapter/gateway"
	"github.com/ksrnnb/chat-app-server/infrastructure/session"
	"github.com/ksrnnb/chat-app-server/usecase"
)

func getMessages(c *gin.Context) {
	// TODO: 認証のミドルウェアをつくる
	s := session.NewSession(c)
	_, ok := s.Get("userId").(int)

	if !ok {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	// TODO: コンテナ管理
	sqlHandler := gateway.NewSqlHandler()
	messageInteractor := usecase.NewMessageInteractor(gateway.NewMessageRepository(sqlHandler))

	roomId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		
	}

	res := messageController.GetMessages(messageInteractor, roomId)

	c.JSON(res.Code, res.Params)
}