package cli

import (
	"errors"
	"fmt"

	"github.com/urfave/cli"
)

type MyFlags struct {
	TextFlag string
}

func StartCli(mf *MyFlags, args []string) error {
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
			Name:  "text, t",
			Usage: "set contents to text file.",
		},
	}

	app.Action = func(ctx *cli.Context) error {
		var name, value string
		name = args[1]
		value = args[2]
		ctx.Set(name, value)

		mf.TextFlag = ctx.String("text")
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
