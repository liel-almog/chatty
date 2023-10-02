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

	return r
}
