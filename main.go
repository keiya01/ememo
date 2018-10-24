package main

import (
	"os"

	"github.com/keiya01/ememo/cli"
)

func main() {
	var cliFlags cli.CliFlags
	cli.StartCli(&cliFlags, os.Args)
}
