package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ksrnnb/chat-app-server/adapter/controller"
	"github.com/ksrnnb/chat-app-server/infrastructure/middleware"
)

var loginController controller.LoginController
var logoutController controller.LogoutController
var userController controller.UserController
var chatRoomController controller.ChatRoomController
var messageController controller.MessageController
var tokenController controller.TokenController

func init() {
	loginController = controller.NewLoginController()
	logoutController = controller.NewLogoutController()
	userController = controller.NewUserController()
	chatRoomController = controller.NewChatRoomController()
	messageController = controller.NewMessageController()
	tokenController = controller.NewTokenController()
}

func SetRoute(r *gin.Engine) *gin.Engine {
	r.GET("/rooms", getRooms)
	r.GET("/self", getUser)
	r.GET("/rooms/:id", getRoom)
	r.GET("/token", getToken)

	csrfRouter := r.Group("/")
	{
		csrfRouter.Use(middleware.Csrf())
		csrfRouter.POST("/login", login)
		csrfRouter.POST("/logout", logout)
		csrfRouter.POST("/rooms/:id", sendMessage)
	}

	return r
}