package route

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ksrnnb/chat-app-server/adapter/gateway"
	"github.com/ksrnnb/chat-app-server/adapter/request"
	"github.com/ksrnnb/chat-app-server/infrastructure/session"
	"github.com/ksrnnb/chat-app-server/usecase"
)

func sendMessage(c *gin.Context) {
	s := session.NewSession(c)
	userId, ok := s.Get("userId").(int)

	if !ok {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	var req request.SendMessageRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: コンテナ管理
	sqlHandler := gateway.NewSqlHandler()
	defer sqlHandler.Close()
	MessageInteractor := usecase.NewMessageInteractor(gateway.NewMessageRepository(sqlHandler))

	roomId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	req.UserId = userId
	req.RoomId = roomId

	res := messageController.SendMessage(req, MessageInteractor)

	c.JSON(res.Code, res.Params)
}