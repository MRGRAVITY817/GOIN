package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/MRGRAVITY817/goin/explorer"
	"github.com/MRGRAVITY817/goin/rest"
)

func usage() {
	fmt.Printf("Welcome to GOIN\n\n")
	fmt.Printf("Plz use the following flags:\n")
	fmt.Printf("-port:   Start the PORT of the server\n")
	fmt.Printf("-mode:   Choose between 'html' and 'rest'\n\n")
	runtime.Goexit() // will terminate this after all the defer options are closed
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}
	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose Between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
}
