package cli

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/keiya01/ememo/file"
	"github.com/keiya01/ememo/format"
	"github.com/keiya01/ememo/input"
)

type TextFlag struct {
	Value string
}

func NewTextFlag(value string) (*TextFlag, error) {
	if value == "" {
		return &TextFlag{}, errors.New("値を空にすることは出来ません")
	}

	return &TextFlag{Value: value}, nil
}

func (t *TextFlag) FlagAction() error {
	fmt.Print("保存するファイル名を入力してください：")
	fileName, err := input.GetUserInputValue()
	if err != nil {
		return err
	}

	t.save(fileName)

	return nil
}

func (t TextFlag) save(fileName string) string {
	setFile := input.AddExtension(fileName)
	fileData, err := os.OpenFile(setFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer fileData.Close()

	contents := format.ChengeToMarkdown(t.Value)

	//書き込み処理
	fmt.Fprintln(fileData, contents)

	fileContents := file.PrintReadFile(setFile)

	log.Printf("TODOを追加しました")
	fmt.Printf(fileContents)
	fmt.Print("=====END=====")

	return fileContents
}
