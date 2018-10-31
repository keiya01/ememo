package main

import (
	"os"

	"github.com/keiya01/ememo/cli"
)

func main() {
	var cf cli.CliFlags
	cli.StartCli(&cf, os.Args)
}
