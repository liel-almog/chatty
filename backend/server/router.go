package server

import (
	"chatty/backend/controllers"
	"chatty/backend/server/websocket"

	"github.com/gin-gonic/gin"
)

var hubMap map[string]*websocket.Hub

func newRouter() *gin.Engine {
	r := gin.Default()
	health := new(controllers.HealthController)

	hubMap = make(map[string]*websocket.Hub)

	r.GET("/health", health.Status)
	r.GET("/ws/chat/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		hub, ok := hubMap[roomId]

		if !ok {
			hub = websocket.NewHub()
			hubMap[roomId] = hub
		}

		websocket.ServeWs(hub, c)
	})

	api := r.Group("/api")
	{
		roomGroup := api.Group("/room")
		{
			roomController := controllers.GetRoomController()
			roomGroup.GET("/", roomController.GetAll)
			roomGroup.GET("/:id/message", roomController.GetAllMessagesByRoom)
		}
	}

	return r
}
