package cmd

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

func CheckingUserInputArgumentValue(args []string) error {
	if len(args) < 3 {
		return errors.New("ERROR: 引数を2つ以上入力してください。")
	}
	return nil
}

func AddExtension(fileName string) string {
	spliteName := strings.Split(fileName, ".")
	setFile := spliteName[0] + ".txt"

	return setFile
}

func GetUserInputValue() (string, error) {
	// Scannerを使って一行読み
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputText := scanner.Text()
	log.Println("input value is ", inputText)

	if inputText == "" {
		return "", errors.New("入力値を空にすることは出来ません")
	}
	return inputText, nil
}
