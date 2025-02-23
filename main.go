package main

import (
	"github.com/murex/gamekit-coffeemachine/cli"
	"os"
)

// main is the entry point of the coffee machine command line runner
func main() {
	cli.Run(os.Args)
}
