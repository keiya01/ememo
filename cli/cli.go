package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

type CmdFlags struct {
	SetFlag string
}

func StartCli(mf *CmdFlags, args []string) error {
	var err error
	err = checkingUserInputValue(args)
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

		mf.SetFlag = ctx.String("set")
		return nil
	}

	err = app.Run(args)
	if err != nil {
		return err
	}

	return nil
}

func checkingUserInputValue(args []string) error {
	if len(args) < 2 {
		return errors.New("ERROR: 引数を入力せずに実行することは出来ません。")
	}
	return nil
}

func (cf CmdFlags) saveInputText(fileName string) string {
	setFile := addExtension(fileName)
	file, err := os.OpenFile(setFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	//書き込み処理
	fmt.Fprintln(file, cf.SetFlag)

	contents := printReadFile(setFile)

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

func addExtension(fileName string) string {
	setFile := fileName
	isTxt := strings.HasSuffix(setFile, ".txt")
	if !isTxt {
		setFile += ".txt"
	}

	return setFile
}
