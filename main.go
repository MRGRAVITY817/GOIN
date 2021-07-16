package main

import (
	"github.com/MRGRAVITY817/goin/cli"
	"github.com/MRGRAVITY817/goin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
