package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ksrnnb/chat-app-server/adapter/controller"
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
	r.GET("/rooms/:id", getMessages)
	return r
}