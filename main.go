package main

import (
	"github.com/MRGRAVITY817/goin/cli"
	"github.com/MRGRAVITY817/goin/db"
)

func main() {
	// defer will activate following function when process has ended
	defer db.Close()
	cli.Start()
}
