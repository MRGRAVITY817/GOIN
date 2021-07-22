package p2p

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type peers struct {
	v map[string]*peer
	m sync.Mutex // Mutex will lock to prevent data races
}

var Peers peers = peers{
	v: make(map[string]*peer),
}

type peer struct {
	key     string
	address string
	port    string
	conn    *websocket.Conn
	inbox   chan []byte
}

func (p *peer) close() {
	Peers.m.Lock()         // mutex.Lock() to start working and block others to access Peers
	defer Peers.m.Unlock() // when finished, allow others to access Peers
	p.conn.Close()
	delete(Peers.v, p.key)
}

func (p *peer) read() {
	// defer will be executed after the function is over.
	defer p.close()
	for {
		m := Message{}
		err := p.conn.ReadJSON(&m)
		if err != nil {
			break
		}
		fmt.Println(m.Kind)
	}
}

func (p *peer) write() {
	defer p.close()
	for {
		m, ok := <-p.inbox // this will block until channel gets message to write
		if !ok {
			break
		}
		p.conn.WriteMessage(websocket.TextMessage, m)
	}
}

// Return all nodes(peers) connected
func AllPeers(p *peers) []string {
	p.m.Lock()
	defer p.m.Unlock()
	var keys []string
	for key := range p.v {
		keys = append(keys, key)
	}
	return keys
}

// Initialize the peer
func initPeer(conn *websocket.Conn, address, port string) *peer {
	key := fmt.Sprintf("%s:%s", address, port)
	p := &peer{
		key:     key,
		address: address,
		port:    port,
		conn:    conn,
		inbox:   make(chan []byte),
	}
	go p.read()
	go p.write()
	Peers.v[key] = p
	return p
}
