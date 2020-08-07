package main

import (
	"../../internal/websocket"
	"github.com/gin-gonic/gin"
)

func main() {
	go websocket.H.Run()

	router := gin.New()
	router.LoadHTMLFiles("../../public/index.html")

	router.GET("/room/:roomId/:name", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/ws/:roomId/:name", func(c *gin.Context) {
		roomId := c.Param("roomId")
		name := c.Param("name")
		websocket.ServeWs(c.Writer, c.Request, roomId, name)
	})

	router.Run("0.0.0.0:8080")
}