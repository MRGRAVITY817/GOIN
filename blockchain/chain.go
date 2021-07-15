package blockchain

import (
	"sync"

	"github.com/MRGRAVITY817/goin/db"
	"github.com/MRGRAVITY817/goin/utils"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once

// It restores byte data to golang data
func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

func (b *blockchain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		// this will continue until reached to genesis block
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() { // make one instance no matter what
			b = &blockchain{"", 0} // Create empty Block chain
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				// create new genesis block
				b.AddBlock("Genesis")
			} else {
				// restore blockchain from bytesfromBytes
				b.restore(checkpoint)
			}
		})
	}
	return b
}
