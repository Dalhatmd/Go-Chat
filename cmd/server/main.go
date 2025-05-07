package main

import (
	"github.com/gin-gonic/gin"
	"github.com/dalhatmd/Go-Chat/internal/websocket"
	"github.com/dalhatmd/Go-Chat/internal/handlers"
)


func main() {
	hub := websocket.NewHub()
	go hub.Run()

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	router.GET("/", func (c *gin.Context) {
		c.HTML(200, "Index.html", gin.H{
			"title": "Homepage",
		})
	})

	router.GET("/ws", handlers.ServeWs(hub))

	router.Run(":8080")
}
