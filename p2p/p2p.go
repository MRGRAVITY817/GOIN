package p2p

import (
	"net/http"

	"github.com/MRGRAVITY817/goin/utils"
	"github.com/gorilla/websocket"
)

// we make connections slice to get multiple user
var conns []*websocket.Conn
var upgrader = websocket.Upgrader{}

// Converting http req,res to websocket method is called "upgrade"
func Upgrade(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	conns = append(conns, conn)
	utils.HandleErr(err)
	for {
		// this will block for reading
		// so the writing have to wait this.
		_, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		// this comes with a problem when the socket disconnects
		for _, aConn := range conns {
			if aConn != conn {
				utils.HandleErr(aConn.WriteMessage(websocket.TextMessage, p))
			}
		}
	}
	conn.Close()
}
