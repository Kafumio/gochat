package router

import (
	"github.com/gin-gonic/gin"
	"gochat/controller"
	"gochat/middleware"
	"gochat/websocket"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	server := gin.Default()
	server.Use(middleware.Cors())
	server.Use(middleware.Recovery)

	socket := websocket.RunSocket

	group := server.Group("")
	{
		group.GET("/user", controller.GetUserList)
		group.GET("/user/:uuid", controller.GetUserDetails)
		group.GET("/user/name", controller.GetUserOrGroupByName)
		group.POST("/user/register", controller.Register)
		group.POST("/user/login", controller.Login)
		group.PUT("/user", controller.ModifyUserInfo)

		group.POST("/friend", controller.AddFriend)

		group.GET("/message", controller.GetMessage)

		group.GET("/file/:fileName", controller.GetFile)
		group.POST("/file", controller.SaveFile)

		group.GET("/group/:uuid", controller.GetGroup)
		group.POST("/group/:uuid", controller.SaveGroup)
		group.POST("/group/join/:userUuid/:groupUuid", controller.JoinGroup)
		group.GET("/group/user/:uuid", controller.GetGroupUsers)

		group.GET("/socket.io", socket)
	}
	return server
}
