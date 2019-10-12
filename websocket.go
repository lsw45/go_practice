package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	// "io/ioutil"
	"log"
	"net/http"
)

/*type message struct {
	userid string
	event  string
}*/

/*func TestXX(t *testing.T) {
	url := "ws://127.0.0.1:8080/ws"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		log.Fatalf("dial websocket url fail: %v", err)
	}

	rm := message{"11", "pushto11"}

	if err := conn.WriteJSON(rm); err != nil {
		log.Fatalf("registe fail: %v", err)
	}

}*/

/*type WsServer struct {
	upgrade *websocket.Upgrader
}

func NewWsServer() *WsServer {
	ws.upgrade = &websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			if r.Method != "GET" {
				fmt.Println("method is not GET")
				return false
			}
			if r.URL.Path != "/ws" {
				fmt.Println("path error")
				return false
			}
			return true
		},
	}
	return ws
}

func (self *WsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := self.upgrade.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("websocket error:", err)
		return
	}
	fmt.Println("client connect :", conn.RemoteAddr())
	go self.connHandle(echo)

}

func echo(conn *websocket.Conn) {

}*/
var upgrader = websocket.Upgrader{} // use default options
func TestXX(t *testing.T) {
	http.HandleFunc("/shiming", Echo)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func Echo(w http.ResponseWriter, r *http.Request) {
	// var conn *websocket.Conn
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	for {
		// var reply string
		mt, message, error := conn.ReadMessage()
		if error != nil {
			fmt.Println("不能够接受消息 error==", error)
			break
		}
		fmt.Println("能够接受到消息了--- ")

		if error = conn.WriteMessage(mt, message); error != nil {
			fmt.Println("不能够发送消息 悲催哦")
			break
		}
	}
}
