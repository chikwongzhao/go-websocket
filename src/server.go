package main

import (
	"fmt"
	"go-websocket/src/impl"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	// 允许跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(rw http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		err    error
		data   []byte
		conn   *impl.Connection
	)
	if wsConn, err = upgrader.Upgrade(rw, r, nil); err != nil {
		fmt.Println(err)
		return
	}
	if conn, err = impl.InitConnection(wsConn); err != nil {
		goto ERR
	}
	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	fmt.Println("listening on 0.0.0.0:7777")
	http.ListenAndServe("0.0.0.0:7777", nil)
}
