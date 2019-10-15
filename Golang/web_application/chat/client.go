package main

import (
	"github.com/gorilla/websocket"
	"time"
)

// clientはチャットを行なっている一人のユーザーを表す
type client struct {
	//  socketはこのクライアントのためのWebSocketです。
	socket *websocket.Conn
	//  sendはメッセージが送られてくるチャンネル
	send chan *message
	// roomはこのクラアントが参加しているチャットルーム
	room *room
	//　userDataはユーザーに関する情報を保持します
	userData map[string]interface{}
}

func (c *client) read() {
	for {
		var msg *message
		if err := c.socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now()
			msg.Time = msg.When.Format("2006/1/2 15:04:05")
			msg.Name = c.userData["name"].(string)
			msg.AvatarURL, _ = c.room.avatar.GetAvatarURL(c)
			if avatarURL, ok := c.userData["avator_url"]; ok {
				msg.AvatarURL = avatarURL.(string)
			}
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteJSON(&msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
