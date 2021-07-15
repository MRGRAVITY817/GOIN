package blockchain

import (
	"sync"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) AddBlock(data string) {
	createBlock(data, b.NewestHash, b.Height)
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() { // make one instance no matter what
			b = &blockchain{"", 0} // Create empty Block chain
			b.AddBlock("Genesis")
		})
	}
	return b
}
