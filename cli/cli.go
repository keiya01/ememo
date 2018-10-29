package cli

import (
	"fmt"
	"os"

	"github.com/keiya01/ememo/cmd"
	"github.com/keiya01/ememo/files"
	"github.com/urfave/cli"
)

//CliFlags is a summary of user input
type CliFlags struct {
	FileNameFlag string
	SetFlag      string
}

func StartCli(cf *CliFlags, args []string) error {
	var err error
	err = cmd.CheckingUserInputValue(args)
	if err != nil {
		fmt.Println(err)
		return err
	}

	app := cli.NewApp()
	app.Name = "ememo"
	app.Usage = "簡単にテキストファイルにメモを作成するツールです。"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "set, s",
			Usage: "set contents to text file.",
		},
	}

	app.Action = func(ctx *cli.Context) error {
		var name, value string
		name = args[1]
		value = args[2]
		ctx.Set(name, value)

		cf.SetFlag = ctx.String("set")
		return nil
	}

	err = app.Run(args)
	if err != nil {
		return err
	}

	return nil
}

func (cf CliFlags) save(fileName string) string {
	setFile := cmd.AddExtension(fileName)
	file, err := os.OpenFile(setFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	//書き込み処理
	fmt.Fprintln(file, cf.SetFlag)

	contents := files.PrintReadFile(setFile)

	fmt.Printf("---- TODOを追加しました ---- \n\n %s", contents)

	return contents
}
