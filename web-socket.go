package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	conn *websocket.Conn
}

var clients = make(map[*Client]bool)
var broadcast = make(chan []byte)

func handleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	//defer conn.Close()

	client := &Client{conn: conn}
	clients[client] = true

	go client.readPump()
}

func (c *Client) readPump() {

	for {
		fmt.Println("two...")
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		broadcast <- msg
	}
}

func handleBroadcast() {
	for {
		message := <-broadcast
		for client := range clients {
			err := client.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				fmt.Println(err)
				client.conn.Close()
				delete(clients, client)
			}
		}
	}
}
