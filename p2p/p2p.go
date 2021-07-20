package p2p

import (
	"fmt"
	"net/http"

	"github.com/MRGRAVITY817/goin/utils"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// Converting http req,res to websocket method is called "upgrade"
func Upgrade(rw http.ResponseWriter, r *http.Request) {
	// Port :3000 will upgrade the request from :4000
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)
	initPeer(conn, "x", "dd")
}

// This will dial(connect) another process to send a message
func AddPeer(address, port string) {
	// Port :4000 is requesting an upgrade from the port :3000
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws", address, port), nil)
	utils.HandleErr(err)
	initPeer(conn, address, port)
}
