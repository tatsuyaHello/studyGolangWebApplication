package main

import (
	"github.com/gorilla/websocket"
)

// clientはチャットを行なっている1人のユーザを表す。
type client struct {
	// socketはこのクライアントのためのWebsocketである。
	socket *websocket.Conn
	// sendはメッセージが送られるチャネルである。
	send chan []byte
	// roomはこのクライアントが参加しているチャットルームである。
	room *room
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			return
		}
	}
}
