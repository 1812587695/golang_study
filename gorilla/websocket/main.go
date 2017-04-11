package main

import (
	//	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")

	})

	http.HandleFunc("/v1/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {
			for {
				mType, msg, _ := conn.ReadMessage()
				conn.WriteMessage(mType, msg)
			}
		}(conn)
	})

	http.HandleFunc("/v2/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {
			for {
				_, msg, _ := conn.ReadMessage()
				println(string(msg)) // string函数是将msg为byte类型的转为字符串
				//				conn.WriteJSON("{'user':'123'}")
				conn.WriteJSON("{'user':'" + string(msg) + "'}")
			}
		}(conn)
	})

	http.HandleFunc("/v3/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {

			ch := time.Tick(5 * time.Second)

			for range ch {
				conn.WriteJSON(myStruct{
					Username:  "msdfsdf",
					FirstName: "234234",
					LastName:  "sdfsldfj",
				})
			}

		}(conn)
	})

	http.HandleFunc("/v4/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil)

		go func(conn *websocket.Conn) {
			for {
				_, msg, _ := conn.ReadMessage() //

				conn.WriteJSON("{'user':'" + string(msg) + "'}")
			}
		}(conn)
	})

	http.ListenAndServe(":3000", nil)
}

type myStruct struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
