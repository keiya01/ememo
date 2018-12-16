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
	var contents []string
	lineNum := 1

	fileContents := file.PrintReadFile(t.Value)
	if fileContents != "" {
		contents = strings.Split(fileContents, "\n")
		for i, content := range contents {
			// 最終行に空文字が含まれているなら削除する
			if i == len(contents)-1 {
				if contents[i] == "" {
					contents = append(contents[:0], contents[:len(contents)-1]...)
					fmt.Print("\n")
					continue
				}
			}
			fmt.Printf("\nememo[line %d]: %s", lineNum, content)
			lineNum++
		}
	} else {
		fmt.Print("======= TODOを入力してください =======\n")
	}

	for {
		fmt.Printf("ememo[line %d]: ", lineNum)
		text := input.GetUserInputValue()

		textType := strings.Split(text, " ")
		if len(textType) > 1 {
			selectedLineNum, err := strconv.Atoi(textType[1])
			if err == nil {
				switch textType[0] {
				case "-e":
					fmt.Printf("ememo_edit[line %d]: ", selectedLineNum)
					newEditTxt := input.GetUserInputValue()
					if newEditTxt == "" {
						continue
					}
					contents[selectedLineNum-1] = newEditTxt
					continue

				case "-d":
					fmt.Printf("delete success line %d \n", selectedLineNum)
					contents = append(contents[:selectedLineNum-1], contents[selectedLineNum:]...)
					continue
				}
			}

		}
		if textType[0] == "-s" {
			fmt.Print("\n")
			for _, content := range contents {
				txt := format.ChangeToMarkdown(content, true)
				fmt.Printf("%s", txt)
			}
			fmt.Print("\n\n")
			continue
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
	fileData, err := os.OpenFile(t.Value, os.O_RDWR|os.O_CREATE, 0666)
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
