package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ksrnnb/chat-app-server/adapter/controller"
)

var loginController controller.LoginController
var logoutController controller.LogoutController
var userController controller.UserController
var chatRoomController controller.ChatRoomController
var messageController controller.MessageController

func init() {
	loginController = controller.NewLoginController()
	logoutController = controller.NewLogoutController()
	userController = controller.NewUserController()
	chatRoomController = controller.NewChatRoomController()
	messageController = controller.NewMessageController()
}

func SetRoute(r *gin.Engine) *gin.Engine {
	r.POST("/login", login)
	r.POST("/logout", logout)
	r.GET("/rooms", getRooms)
	r.GET("/self", getUser)
	r.GET("/rooms/:id", getRoom)
	r.POST("/rooms/:id", sendMessage)
	return r
}