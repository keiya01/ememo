package cli

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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
	var contents []string
	lineNum := 1
	for {
		fmt.Print("======= TODOを入力してください =======\n")
		fmt.Printf("ememo[line %d]: ", lineNum)
		text, err := input.GetUserInputValue()
		if err != nil {
			return err
		}

		edit := strings.Split(text, " ")
		if edit[0] == "edit" {
			editLineNum, err := strconv.Atoi(edit[1])
			if err == nil {
				fmt.Printf("ememo edit(line %d): ", editLineNum)
				newEditTxt, err := input.GetUserInputValue()
				if err != nil {
					return err
				}

				contents[editLineNum-1] = newEditTxt

				continue
			}
		}

		if text == "end" {
			break
		}

		contents = append(contents, text)

		lineNum++
	}

	t.save(contents)

	return nil
}

func (t TextFlag) save(contents []string) string {
	fileData, err := os.OpenFile(t.Value, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		color.Red("ERROR: %v", err)
		return ""
	}
	defer fileData.Close()

	for _, text := range contents {
		contents := format.ChengeToMarkdown(text)
		//書き込み処理
		fmt.Fprintln(fileData, contents)
	}

	fileContents := file.PrintReadFile(t.Value)

	log.Printf("TODOを追加しました")

	return fileContents
}
