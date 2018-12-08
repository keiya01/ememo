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

func (t *TextFlag) FlagAction() {
	fileContents := file.PrintReadFile(t.Value)
	if fileContents == "" {
		fmt.Printf("\n%s\n", fileContents)
	}

	var contents []string
	lineNum := 1
	fmt.Print("======= TODOを入力してください =======\n")
	for {
		fmt.Printf("ememo[line %d]: ", lineNum)
		text := input.GetUserInputValue()

		edit := strings.Split(text, " ")
		if edit[0] == "edit" {
			editLineNum, err := strconv.Atoi(edit[1])
			if err == nil {
				fmt.Printf("ememo_edit[line %d]: ", editLineNum)
				newEditTxt := input.GetUserInputValue()
				if newEditTxt == "" {
					continue
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

	return
}

func (t TextFlag) save(contents []string) string {
	fileData, err := os.OpenFile(t.Value, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		color.Red("ERROR: %v", err)
		return ""
	}
	defer fileData.Close()

	var fileContents string
	totalContants := len(contents) - 1
	for i, text := range contents {
		fileContents += format.ChangeToMarkdown(text, i < totalContants)
	}

	//書き込み処理
	fmt.Fprintln(fileData, fileContents)

	body := file.PrintReadFile(t.Value)

	log.Printf("TODOを追加しました")

	return body
}
