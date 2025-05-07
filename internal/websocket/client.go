package websocket

import (
	"github.com/gorilla/websocket"
	"fmt"
)


type Client struct {
	Conn *websocket.Conn
	Send chan[]byte
	Hub *Hub
}

func (c *Client) ReadPump() {
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Print(err)
			c.Hub.Unregister <- c
			c.Conn.Close()
			break
		}
		c.Hub.Broadcast <- message
	}
}

func (c *Client) WritePump() {
	defer c.Conn.Close()

	for message := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			fmt.Print(err)
			break
		}
	}
}

