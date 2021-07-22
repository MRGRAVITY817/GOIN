package p2p

import (
	"fmt"
	"net/http"
	"time"

	"github.com/MRGRAVITY817/goin/utils"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// Converting http req,res to websocket method is called "upgrade"
func Upgrade(rw http.ResponseWriter, r *http.Request) {
	openPort := r.URL.Query().Get("openPort")
	ip := utils.Splitter(r.RemoteAddr, ":", 0)
	// Port :3000 will upgrade the request from :4000
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return openPort != "" && ip != ""
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)
	peer := initPeer(conn, ip, openPort)
	time.Sleep(2 * time.Second)
	peer.inbox <- []byte("Hello from 3000!")
}

// This will dial(connect) another process to send a message
func AddPeer(address, port, openPort string) {
	// Port :4000 is requesting an upgrade from the port :3000
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws?openPort=%s", address, port, openPort[1:]), nil)
	utils.HandleErr(err)
	peer := initPeer(conn, address, openPort)
	time.Sleep(1 * time.Second)
	peer.inbox <- []byte("Hello from 4000!")
}
