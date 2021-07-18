package blockchain

import (
	"sync"

	"github.com/MRGRAVITY817/goin/db"
	"github.com/MRGRAVITY817/goin/utils"
)

type blockchain struct {
	NewestHash        string `json:"newestHash"`
	Height            int    `json:"height"`
	CurrentDifficulty int    `json:"currentDifficulty"`
}

const (
	defaultDifficulty  int = 2
	difficultyInterval int = 5
	blockInterval      int = 2
	allowedRange       int = 2
)

var b *blockchain
var once sync.Once

// It restores byte data to golang data
func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) AddBlock() {
	block := createBlock(b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	b.persist()
}

// Check if the block mining interval is in the range
// of the time length that we expected to be in.
func (b *blockchain) recalculateDifficulty() int {
	allBlocks := b.Blocks()
	newestBlock := allBlocks[0]
	lastRecalculatedBlock := allBlocks[difficultyInterval-1]
	actualTime := (newestBlock.Timestamp / 60) - (lastRecalculatedBlock.Timestamp / 60)
	expectedTime := difficultyInterval * blockInterval
	if actualTime <= (expectedTime - allowedRange) {
		return b.CurrentDifficulty + 1
	} else if actualTime > (expectedTime + allowedRange) {
		return b.CurrentDifficulty - 1
	}
	return b.CurrentDifficulty
}

func (b *blockchain) difficulty() int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		return b.recalculateDifficulty()
	}
	return b.CurrentDifficulty
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
	db.SaveCheckpoint(utils.ToBytes(b))
}

func (b *blockchain) UTxOutsByAddress(address string) []*UTxOut {
	var uTxOuts []*UTxOut
	creatorTxs := make(map[string]bool)
	for _, block := range b.Blocks() {
		for _, tx := range block.Transactions {
			// Saving the txId that created the output
			// that user is using as an input.
			for _, input := range tx.TxIns {
				if input.Owner == address {
					creatorTxs[input.TxId] = true
				}
			}
			// We won't use the already-spent txs as input
			for index, output := range tx.TxOuts {
				if output.Owner == address {
					if _, ok := creatorTxs[tx.Id]; !ok {
						uTxOuts = append(uTxOuts, &UTxOut{tx.Id, index, output.Amount})
					}
				}
			}
		}
	}
	return uTxOuts
}

func (b *blockchain) BalanceByAddress(address string) int {
	txOuts := b.UTxOutsByAddress(address)
	var amount int
	for _, txOut := range txOuts {
		amount += txOut.Amount
	}
	return amount
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() { // make one instance no matter what
			b = &blockchain{Height: 0} // Create empty Block chain
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				// create new genesis block
				b.AddBlock()
			} else {
				// restore blockchain from bytesfromBytes
				b.restore(checkpoint)
			}
		})
	}
	return b
}
