package main

import (
	"os"

	"github.com/keiya01/ememo/cli"
)

func main() {
	var cmdFlags cli.CmdFlags
	cli.StartCli(&cmdFlags, os.Args)
}
