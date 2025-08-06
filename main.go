package main

import (
	"os"

	"github.com/murex/gamekit-coffeemachine/cli"
)

// main is the entry point of the coffee machine command line runner
func main() {
	cli.Run(os.Args)
}
