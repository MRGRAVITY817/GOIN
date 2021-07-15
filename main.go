package main

import (
	"github.com/MRGRAVITY817/goin/blockchain"
)

func main() {
	blockchain.Blockchain().AddBlock("First")
	blockchain.Blockchain().AddBlock("Second")
	blockchain.Blockchain().AddBlock("Third")
	blockchain.Blockchain().AddBlock("Fourth")
	blockchain.Blockchain().AddBlock("Five")
}
