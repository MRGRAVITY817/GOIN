package p2p

import (
	"fmt"
	"net/http"

	"github.com/MRGRAVITY817/goin/utils"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)
	for {
		_, p, err := conn.ReadMessage()
		utils.HandleErr(err)
		fmt.Printf("%s\n", p)
	}
}
