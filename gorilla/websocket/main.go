package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type Hub struct {
	clients   map[*Client]bool
	addClient chan *Client
	msg       chan []byte
}

var hub = Hub{
	clients:   make(map[*Client]bool),
	addClient: make(chan *Client),
	msg:       make(chan []byte),
}

func (hub *Hub) start() {
	fmt.Println("123")
	for {
		select {
		// 读取入channel中的数据（1channel）
		case conn := <-hub.addClient:
			hub.clients[conn] = true // 保存map

			fmt.Println(hub.clients)
		// 读取channel中的数据，这里是获取用户发送的数据（2channel）
		case msg := <-hub.msg:

			// 这里循环map里面保存的客户端连接信息
			for k, _ := range hub.clients {
				//发送消息到每一个连接的客户端
				k.ws.WriteMessage(1, msg)
			}
		}
	}
}

type Client struct {
	ws *websocket.Conn
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")

	})

	// 首先main函数走到这里，将数据保存全局变量里面
	go hub.start()

	http.HandleFunc("/v1/ws", func(w http.ResponseWriter, r *http.Request) {

		// 创建客户端连接
		var conn, _ = upgrader.Upgrade(w, r, nil)

		client := &Client{ws: conn}

		// 将client写入channel中（1channe）
		hub.addClient <- client

		// 这里获取客户端发送来的消息
		go func(conn *websocket.Conn) {
			for {
				_, msg, _ := conn.ReadMessage()

				//将msg保存在channel中（2channel）
				hub.msg <- msg
			}
		}(conn)
	})

	http.ListenAndServe(":3000", nil)
}
