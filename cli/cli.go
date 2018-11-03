package cli

import (
	"errors"

	"github.com/fatih/color"
	"github.com/keiya01/ememo/format"
	"github.com/keiya01/ememo/input"
	"github.com/urfave/cli"
)

func StartCli(args []string) error {
	err := input.CheckingUserInputArgumentValue(args)
	if err != nil {
		color.Red("ERROR: %v", err)
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
		cli.BoolFlag{
			Name:  "mark, m",
			Usage: "show markdown contents.",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.Bool("mark") {
			format.ShowMarkdown()
			return nil
		}

		textFlag, err := NewTextFlag(c.String("text"))
		if err != nil {
			color.Red("ERROR: %v", err)
			return err
		}
		err = textFlag.FlagAction()
		if err != nil {
			color.Red("ERROR: %v", err)
			return err
		}

		err = errors.New("「" + args[1] + "」オプションは使用できません。「-h」オプションで確認してください。")
		return err
	}

	app.Run(args)

	return nil
}
