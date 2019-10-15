package main

import (
	"github.com/RyotaNakajimakun/Training/Golang/web_application/trace"
	"github.com/gorilla/websocket"
	"github.com/stretchr/objx"
	"log"
	"net/http"
)

type room struct {
	// forwardは他のクライアントに転送するためのメッセージを保持するチャンネル
	forward chan *message

	// joinはチャットルームに参加しようとしているクライアントのためのチャンネル
	join chan *client

	// leaveはチャットルームから退室しようとしているクライアントのためのチャンネル
	leave chan *client

	// clientには在室している全てのクライアントが保存される
	clients map[*client]bool

	// tracerはチャットルーム上で行われた操作のログを受け取ります。
	tracer trace.Tracer

	// avatarはアバターの情報を取得します
	avatar Avatar
}

func newRoom(avatar Avatar) *room {
	return &room{
		forward: make(chan *message),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		tracer:  trace.Off(),
		avatar:  avatar,
	}
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			//　参加
			r.clients[client] = true
			r.tracer.Trace("新しいクライアントが参加しました")

		case client := <-r.leave:
			// 退室
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("クライアントが退室しました")

		case msg := <-r.forward:
			r.tracer.Trace("メッセージを受信しました: ", msg.Message)
			for client := range r.clients {
				select {
				// メッセージを送信
				case client.send <- msg:
					r.tracer.Trace("メッセージを送信しました: ", msg.Message)

				default:
					// 送信に失敗
					delete(r.clients, client)
					close(client.send)
					r.tracer.Trace(" -- 送信に失敗しました。クライアントをクリーンアップします")
				}
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: messageBufferSize,
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServerHTTP:", err)
		return
	}

	authCookie, err := req.Cookie("auth")
	if err != nil {
		log.Fatal("クッキーの取得に失敗しました:", err)
		return
	}
	client := &client{
		socket: socket,
		send:   make(chan *message, messageBufferSize),
		room:   r,
		userData: objx.MustFromBase64(authCookie.Value),
	}

	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}
