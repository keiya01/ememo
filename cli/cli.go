package cli

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/keiya01/ememo/file"
	"github.com/keiya01/ememo/format"
	"github.com/keiya01/ememo/input"
	"github.com/urfave/cli"
)

type flags interface {
	FlagAction() error
}

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
			Usage: "set contents to text file.\nPlease put the text in the first argument.",
		},
		cli.StringFlag{
			Name:  "read, r",
			Usage: "show text file contents.\nPlease put the file name in the first argument.",
		},
		cli.StringFlag{
			Name:  "comp, c",
			Usage: "Check the completed TODO.\nPlease enter the file name as an argument.",
		},
		cli.BoolFlag{
			Name:  "mark, m",
			Usage: "show markdown contents.\nPlease do not enter anything in the argument.",
		},
	}

	app.Action = func(c *cli.Context) error {

		if text := c.String("text"); text != "" {
			textFlag := NewTextFlag(text)
			err = textFlag.FlagAction()
			if err != nil {
				color.Red("ERROR: %v", err)
				return err
			}
			return nil
		}

		if read := c.String("read"); read != "" {
			setFile := input.AddExtension(read)
			fileContents := file.PrintReadFile(setFile)
			fmt.Print("===== TODO LIST =====")
			color.Blue("\n%s", fileContents)
			fmt.Print("===== END =====")
			return nil
		}

		if comp := c.String("comp"); comp != "" {
			cf := NewCompFlag(comp)
			_, err := cf.FlagAction()
			if err != nil {
				color.Red("ERROR: %v", err)
				return err
			}
			return nil
		}

		if c.Bool("mark") {
			format.ShowMarkdown()
			return nil
		}
		err = errors.New("「" + args[1] + "」オプションは使用できません。「-h」オプションで確認してください。")
		color.Red("ERROR: %v", err)
		return err
	}

	app.Run(args)

	return nil
}
