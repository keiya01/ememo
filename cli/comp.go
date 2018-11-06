package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/keiya01/ememo/file"
	"github.com/keiya01/ememo/input"
)

type CompFlag struct {
	FileName string
}

func NewCompFlag(fileName string) *CompFlag {
	setFile := input.AddExtension(fileName)

	return &CompFlag{FileName: setFile}
}

func (c *CompFlag) FlagAction() (string, error) {
	fileData, err := os.OpenFile(c.FileName, os.O_RDWR, 0666)
	if err != nil {
		return "", err
	}
	defer fileData.Close()

	readFileData, err := os.Open(c.FileName)
	if err != nil {
		return "", nil
	}
	defer readFileData.Close()

	todoList := file.FileScan(readFileData,
		func(scanner *bufio.Scanner, index int) string {
			return strconv.Itoa(index) + ". " + scanner.Text() + "\n"
		},
	)
	fmt.Println(todoList)
	fmt.Print("完了したTODOの番号を上記から選んでください: ")
	input, err := input.GetUserInputValue()
	if err != nil {
		return "", err
	}

	sentenceNum, err := strconv.Atoi(input)
	if err != nil {
		return "", err
	}

	contents := c.printReadCompTodo(sentenceNum)
	fmt.Fprintln(fileData, contents)

	fileContents := file.PrintReadFile(c.FileName)

	return fileContents, nil
}

func (c CompFlag) printReadCompTodo(sentenceNum int) string {
	fileData, err := os.Open(c.FileName)
	if err != nil {
		color.Red("ERROR: %v", err)
		return ""
	}
	defer fileData.Close()

	return file.FileScan(fileData, func(scanner *bufio.Scanner, index int) string {
		var compTodo string
		if sentenceNum == index {
			compTodo = strings.Replace(scanner.Text(), "[ ]", "[x]", 1)
			return compTodo + "\n"
		}

		return scanner.Text() + "\n"

	})
}
