package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	ws "github.com/dalhatmd/Go-Chat/internal/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func ServeWs(hub *ws.Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		client := &ws.Client{
			Conn: conn,
			Send: make(chan []byte, 256),
			Hub: hub,
		}
		hub.Register <- client
		
		go client.WritePump()
		go client.ReadPump()
	}
}
