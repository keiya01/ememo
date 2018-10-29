package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/keiya01/ememo/cmd"
	"github.com/urfave/cli"
)

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

func (cf CliFlags) saveInputText(fileName string) string {
	setFile := cmd.AddExtension(fileName)
	file, err := os.OpenFile(setFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	//書き込み処理
	fmt.Fprintln(file, cf.SetFlag)

	contents := printReadFile(setFile)

	fmt.Printf("TODOを追加しました。 \n %s", contents)

	return contents
}

func printReadFile(fileName string) string {
	var contents string

	// ファイルを読み出し用にオープン
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	// 一行ずつ読み出し
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents += scanner.Text()
	}

	return contents
}
