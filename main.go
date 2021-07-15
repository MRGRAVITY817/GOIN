package main

import (
	"github.com/MRGRAVITY817/goin/explorer"
	"github.com/MRGRAVITY817/goin/rest"
)

func main() {
	go explorer.Start(3421)
	rest.Start(4321)
}
