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
var chatRoomController controller.ChatRoomController

func init() {
	loginController = controller.NewLoginController()
	chatRoomController = controller.NewChatRoomController()
}

func SetRoute(r *gin.Engine) *gin.Engine {
	r.POST("/login", login)
	r.GET("/rooms", getRooms)
	r.GET("/self", getUser)
	return r
}

func login(c *gin.Context) {
	var req request.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: コンテナ管理
	interactor := usecase.NewUserInteractor(gateway.NewUserRepository(gateway.NewSqlHandler()))
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

func getUser(c *gin.Context) {
	interactor := usecase.NewUserInteractor(gateway.NewUserRepository(gateway.NewSqlHandler()))
	res := loginController.GetUser(session.NewSession(c), interactor)

	c.JSON(res.Code, res.Params)
}