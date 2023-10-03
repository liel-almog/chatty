package server

import (
	"chatty/backend/controllers"
	"chatty/backend/server/websocket"

	"github.com/gin-gonic/gin"
)

func newRouter() *gin.Engine {
	r := gin.Default()
	health := new(controllers.HealthController)
	hub := websocket.NewHub()

	r.GET("/health", health.Status)
	r.GET("/ws/chat", func(c *gin.Context) {
		websocket.ServeWs(hub, c)
	})

	api := r.Group("/api")
	{
		chatGroup := api.Group("/chat")
		{
			messageGroup := chatGroup.Group("/message")
			{
				messageController := controllers.GetMessageController()
				messageGroup.GET("/", messageController.GetAll)
			}
		}

		roomGroup := api.Group("/room")
		{
			roomController := controllers.GetRoomController()
			roomGroup.GET("/", roomController.GetAll)
		}
	}

	return r
}
