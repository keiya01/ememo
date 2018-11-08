package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/keiya01/ememo/file"
	"github.com/keiya01/ememo/format"
	"github.com/keiya01/ememo/input"
)

type TextFlag struct {
	Value string
}

func NewTextFlag(value string) *TextFlag {
	file := input.AddExtension(value)
	return &TextFlag{Value: file}
}

func (t *TextFlag) FlagAction() error {
	fmt.Print("TODOを入力してください：")
	text, err := input.GetUserInputValue()
	if err != nil {
		return err
	}

	t.save(text)

	return nil
}

func (t TextFlag) save(text string) string {
	fileData, err := os.OpenFile(t.Value, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		color.Red("ERROR: %v", err)
		return ""
	}
	defer fileData.Close()

	contents := format.ChengeToMarkdown(text)

	//書き込み処理
	fmt.Fprintln(fileData, contents)

	fileContents := file.PrintReadFile(t.Value)

	log.Printf("TODOを追加しました")

	return fileContents
}
