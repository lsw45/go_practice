package test

import (
	"github.com/gorilla/websocket"
	"log"
	"testing"
)

type message struct {
	userid string
	event  string
}

func TestWS(t *testing.T) {
	url := "ws://127.0.0.1:8888/ws"
	conn, resp, err := websocket.DefaultDialer.Dial(url, nil)
	log.Printf("%+v", resp)
	if err != nil {
		log.Fatalf("dial websocket url fail: %v", err)
	}

	rm := message{"11", "pushto11"}

	if err := conn.WriteJSON(rm); err != nil {
		log.Fatalf("registe fail: %v", err)
	}
}
