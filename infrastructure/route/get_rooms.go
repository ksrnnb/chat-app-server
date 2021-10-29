package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ksrnnb/chat-app-server/adapter/gateway"
	"github.com/ksrnnb/chat-app-server/infrastructure/session"
	"github.com/ksrnnb/chat-app-server/usecase"
)

func getRooms(c *gin.Context) {
	s := session.NewSession(c)
	userId, ok := s.Get("userId").(int)

	if !ok {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	// TODO: コンテナ管理
	sqlHandler := gateway.NewSqlHandler()
	chatRoomInteractor := usecase.NewChatRoomInteractor(gateway.NewChatRoomRepository(sqlHandler))

	res := chatRoomController.GetChatRooms(chatRoomInteractor, userId)

	c.JSON(res.Code, res.Params)
}