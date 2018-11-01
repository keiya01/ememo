package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/keiya01/ememo/cmd"
	"github.com/keiya01/ememo/file"
	"github.com/keiya01/ememo/format"
	"github.com/urfave/cli"
)

//CliFlags is a summary of user input
type CliFlags struct {
	FileNameFlag string
	SetFlag      string
}

func StartCli(cf *CliFlags, args []string) error {
	var err error
	err = cmd.CheckingUserInputArgumentValue(args)
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

		fmt.Print("保存するファイル名を入力してください：")
		fileName, err := cmd.GetUserInputValue()
		if err != nil {
			color.Red("Error: %v", err)
			return err
		}

		cf.save(fileName)
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
	contents := format.ChengeToMarkdown(cf.SetFlag)
	//書き込み処理
	fmt.Fprintln(file, contents)

	fileContents := files.PrintReadFile(setFile)

	log.Printf("TODOを追加しました")
	fmt.Printf(fileContents)
	fmt.Print("=====END=====")

	return contents
}
